package config

import (
	"ecommerce/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	Type     string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var dbConfig DBConfig

func InitDB(config DBConfig) *gorm.DB {
	dbConfig = config

	var (
		dialect gorm.Dialector
		dsn     string
		err     error
	)

	switch config.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)
		dialect = mysql.Open(dsn)
	case "pgsql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.Username,
			config.Password,
			config.Database,
		)
		dialect = postgres.Open(dsn)
	default:
		helper.LogError(fmt.Errorf("Unknown database type: %s", config.Type))
	}

	DB, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		helper.LogError(err)
	}
	return DB
}

func GetDBConfig() DBConfig {
	return dbConfig
}
