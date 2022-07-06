package database

import (
	"errors"
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "time"
)

func New() (*gorm.DB, error){

	err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
      }

	host := os.Getenv("DB_LOCALHOST")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_DBNAME")

	s := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, dbName)
	
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{})

	if err != nil {
		return nil, errors.New("gagal konek db")
	  }

	return db, nil
}