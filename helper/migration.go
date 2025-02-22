package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func RunMigrations(db *sql.DB, folderPath string) {
	files, err := filepath.Glob(filepath.Join(folderPath, "*.sql"))
	if err != nil {
		log.Fatalf("Error reading migration files: %v", err)
	}

	for _, file := range files {
		fmt.Printf("Running: %s\n", file)

		// Gunakan os.ReadFile sebagai pengganti ioutil.ReadFile
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}

		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("Error executing %s: %v", file, err)
		}
	}
	fmt.Println("All migrations executed successfully.")
}
