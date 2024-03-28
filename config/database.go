package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shuklarituparn/VK-Marketplace/api/models"
	"github.com/shuklarituparn/VK-Marketplace/internal/logger"

	"github.com/charmbracelet/log"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	return db
}

func ConnectDb() {
	var fileLogger = logger.SetupLogger()

	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)
	portInt, _ := strconv.Atoi(port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)

	postgresqlDb, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Error("Error connecting to the database:", err)
		fileLogger.Println("Error connecting to the database:", err)
	}
	migrationErr := postgresqlDb.AutoMigrate(&models.User{}, &models.Ad{})
	if migrationErr != nil {
		log.Error(migrationErr)
		fileLogger.Println(migrationErr)
	}
	db = postgresqlDb
	log.Info("Successfully connected!")
	fileLogger.Println("Successfully connected to the database!")
}
