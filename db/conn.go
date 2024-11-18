package db

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/laysaalves/url-miuda/internal/entity"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := godotenv.Load("configs/.env")
	if err != nil {
		logger.Fatal("Error loading .env file:", zap.Error(err))
	}
	postgresConn := os.Getenv("POSTGRES_CONNECTION")
	if postgresConn == "" {
		logger.Fatal("POSTGRES_CONNECTION not found in environment variables")
	}
	dsn := postgresConn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	} else if err := db.AutoMigrate(&entity.Link{}); err != nil {
		logger.Fatal("Failed to migrate database", zap.Error(err))
	} else {
		logger.Info("Database created successfully")
	}
}