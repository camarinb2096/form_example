package db

import (
	"camarinb2096/form_example/internal/app/complaint"
	"camarinb2096/form_example/internal/app/customer"
	"camarinb2096/form_example/pkg/logger"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Database string `envconfig:"DB_DATABASE" default:"mydb"`
}

func NewConfig() *Config {

	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

func NewDb(cfg *Config, logger *logger.Logger) *gorm.DB {
	logger.Info("Connecting to database...")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Connected to database: %s ", cfg.Database)

	logger.Info("Migrating database...")
	db.Debug().AutoMigrate(&complaint.Complaint{})
	db.Debug().AutoMigrate(&customer.Customer{})

	logger.Info("Database migrated")
	return db
}

func CloseDb(db *gorm.DB, logger *logger.Logger) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	logger.Info("Closing database connection...")
	return sqlDb.Close()
}
