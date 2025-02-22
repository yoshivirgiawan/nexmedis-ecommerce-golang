package cmd

import (
	"ecommerce/app/modules/product"
	"ecommerce/app/modules/user"
	"ecommerce/config"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeders",
	Run: func(cmd *cobra.Command, args []string) {
		// Ambil konfigurasi database
		dbConfig := config.GetDBConfig()

		// Inisialisasi koneksi database menggunakan GORM
		db := config.InitDB(dbConfig)

		fmt.Println("Running seeders...")

		// Seed data untuk tabel users
		users := []user.User{
			{Name: "admin", Email: "admin@example.com", Password: "password123", PhoneNumber: "081234567890", Role: "admin"},
			{Name: "user1", Email: "user1@example.com", Password: "password123", PhoneNumber: "081234567891"},
		}

		for _, u := range users {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
			if err != nil {
				log.Fatalf("Failed to hash password: %v", err)
			}
			u.Password = string(hashedPassword)
			if err := db.FirstOrCreate(&u, user.User{Email: u.Email}).Error; err != nil {
				log.Fatalf("Failed to seed users: %v", err)
			}
		}
		fmt.Println("Users seeded successfully.")

		imageURLs := []string{
			"https://picsum.photos/id/1/200/300",
			"https://picsum.photos/id/2/200/300",
			"https://picsum.photos/id/3/200/300",
		}

		// Seed data untuk tabel products
		products := []product.Product{
			{Name: "Product 1", Description: "Description for product 1", Price: 10000},
			{Name: "Product 2", Description: "Description for product 2", Price: 20000},
		}

		for _, p := range products {
			rand.Seed(time.Now().UnixNano())
			selectedImageURL := imageURLs[rand.Intn(len(imageURLs))]

			if _, err := os.Stat("storage/public"); os.IsNotExist(err) {
				// Jika tidak ada, buat direktori beserta sub-direktorinya
				err := os.MkdirAll("storage/public", os.ModePerm)
				if err != nil {
					log.Fatalf("Failed to create storage directory: %v", err)
				}
			}
			// Tentukan nama file gambar dan path untuk menyimpannya
			imageFileName := fmt.Sprintf("%d.jpg", rand.Int())
			imagePath := filepath.Join("storage", "public", imageFileName)

			// Unduh dan simpan gambar
			err := downloadImage(selectedImageURL, imagePath)
			if err != nil {
				log.Fatalf("Failed to download and save image: %v", err)
			}
			fmt.Printf("Image downloaded and saved to %s\n", imagePath)

			p.Image = imageFileName
			if err := db.FirstOrCreate(&p, product.Product{Name: p.Name}).Error; err != nil {
				log.Fatalf("Failed to seed products: %v", err)
			}
		}
		fmt.Println("Products seeded successfully.")

		fmt.Println("Seeders completed.")
	},
}

// Fungsi untuk mengunduh gambar dari URL dan menyimpannya ke disk
func downloadImage(url string, savePath string) error {
	// Ambil gambar dari URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// Buat file baru untuk menyimpan gambar
	out, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create image file: %v", err)
	}
	defer out.Close()

	// Salin konten gambar ke file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
