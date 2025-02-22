package cmd

import (
	"ecommerce/app/modules/cart"
	"ecommerce/app/modules/order"
	"ecommerce/app/modules/product"
	"ecommerce/app/modules/user"
	"ecommerce/config"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var fresh bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Ambil konfigurasi database
		dbConfig := config.GetDBConfig()

		// Inisialisasi koneksi database menggunakan GORM
		db := config.InitDB(dbConfig)

		if fresh {
			fmt.Println("Dropping all tables...")
			err := db.Migrator().DropTable(
				&user.User{},
				&product.Product{},
				&cart.Cart{},
				&order.Order{},
				&order.OrderDetail{},
			)
			if err != nil {
				log.Fatalf("Failed to drop tables: %v", err)
			}
			fmt.Println("All tables dropped successfully.")
		}

		fmt.Println("Running migrations...")
		err := db.AutoMigrate(
			&user.User{}, // Entity User
			&product.Product{},
			&cart.Cart{},
			&order.Order{},
			&order.OrderDetail{},
		)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		fmt.Println("Migrations completed successfully.")
	},
}

func init() {
	// Tambahkan flag untuk fresh
	migrateCmd.Flags().BoolVarP(&fresh, "fresh", "f", false, "Drop all tables before migrating")
	rootCmd.AddCommand(migrateCmd)
}
