package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LogError(loggedErr error) {
	// Tentukan direktori untuk menyimpan log berdasarkan tanggal saat ini (format: YYYY-MM-DD)
	logDir := filepath.Join("storage", "logs", "errors")
	// Pastikan direktori untuk menyimpan log ada, jika belum buat direktori tersebut
	err := os.MkdirAll(logDir, 0755)
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
	logLine := fmt.Sprintf("[%s] :\n ERROR: %s\n", time.Now().Format(time.RFC3339), loggedErr.Error())
	_, err = logFile.WriteString(logLine)
	if err != nil {
		fmt.Println("Error writing to log file:", err.Error())
	}
}
