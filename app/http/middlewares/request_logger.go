package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Baca body dari request
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading request body:", err.Error())
		}

		// Salin kembali body yang sudah dibaca agar bisa digunakan lagi
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Buat buffer untuk menangkap respons
		responseBuf := &bytes.Buffer{}
		// Ganti writer ResponseWriter dengan buffer agar bisa menangkap respons
		c.Writer = &responseWriter{body: responseBuf, ResponseWriter: c.Writer}

		start := time.Now()

		// Lanjutkan penanganan request
		c.Next()

		// Tentukan direktori untuk menyimpan log berdasarkan tanggal saat ini (format: YYYY-MM-DD)
		logDir := filepath.Join("storage", "logs", "api")
		// Pastikan direktori untuk menyimpan log ada, jika belum buat direktori tersebut
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			fmt.Println("Error creating log directory:", err.Error())
		}

		// Buat nama file log menggunakan format waktu (format: YYYY-MM-DD-HH-MM-SS.log)
		logFileName := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")

		// Tulis log ke file
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening log file:", err.Error())
		}
		defer logFile.Close()

		// Tulis log ke file
		logLine := fmt.Sprintf("[%s] %s - %s :\n URL: %s\n Method: %s\n Headers: %v\n Body: %s\n Response: %s\n Duration: %v\n", time.Now().Format(time.RFC3339), c.ClientIP(), c.Request.UserAgent(), c.Request.URL.String(), c.Request.Method, c.Request.Header, string(body), responseBuf.String(), time.Since(start))
		_, err = logFile.WriteString(logLine)
		if err != nil {
			fmt.Println("Error writing to log file:", err.Error())
		}
	}
}

// Custom ResponseWriter untuk menangkap respons
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
