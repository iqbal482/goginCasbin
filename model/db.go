package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func DBConnection() (*gorm.DB, error) {
	DB_HOST := "postgresql.postgre.svc.cluster.local"
	DB_PORT := "5432"
	DB_USER := "postgres"
	DB_PASS := "q6PENVTLBW"
	DB_NAME := "postgres"

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