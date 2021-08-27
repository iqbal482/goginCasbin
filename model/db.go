package model

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"time"
)

func goDotEnvVariable(key string) string {

	// load .env file
	//err := godotenv.Load(".env")
	//err, nil := os.Getenv("DB_HOST")
	path, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(path, ".env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func DBConnection() (*gorm.DB, error) {
	//os.Setenv("DB_HOST", "postgresql.postgre.svc.cluster.local")
	//os.Setenv("DB_PORT", "5432")
	//os.Setenv("DB_USER", "postgres")
	//os.Setenv("DB_PASS", "postgres")
	//os.Setenv("DB_NAME", "gogin")

	DB_HOST := goDotEnvVariable("DB_HOST")
	DB_PORT := goDotEnvVariable("DB_PORT")
	DB_USER := goDotEnvVariable("DB_USER")
	DB_PASS := goDotEnvVariable("DB_PASS")
	DB_NAME := goDotEnvVariable("DB_NAME")

	//USER := "postgres"
	//PASS := "postgres"
	//HOST := "127.0.0.1"
	//PORT := 5632
	//DBNAME := "gogin"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		})

	url := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})
}