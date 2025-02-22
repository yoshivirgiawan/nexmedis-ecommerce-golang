package main

import (
	"ecommerce/app/modules/jwtgenerator"
	"ecommerce/cmd"
	"ecommerce/config"
	"ecommerce/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	appPort := "8080"

	if os.Getenv("APP_PORT") != "" {
		appPort = os.Getenv("APP_PORT")
	}

	if os.Getenv("JWT_SECRET_KEY") == "" {
		jwtgenerator.GenerateAndWriteSecretKey()
	}

	dbConfig := config.DBConfig{
		Type:     os.Getenv("DB_CONNECTION"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	redisConfig := config.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	config.InitDB(dbConfig)
	config.InitRedis(redisConfig)

	cmd.Execute()

	r := routes.SetupRouter()
	r.Run(":" + appPort)
}
