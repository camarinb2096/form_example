package db

import (
	"camarinb2096/form_example/internal/app/complaint"
	"camarinb2096/form_example/internal/app/customer"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
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

func NewDb(cfg *Config) *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database:", cfg.Database)

	db.Debug().AutoMigrate(&complaint.Complaint{})
	db.Debug().AutoMigrate(&customer.Customer{})

	log.Println("Database migrated")
	return db
}

func CloseDb(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	log.Println("Closing database connection")
	return sqlDb.Close()
}
