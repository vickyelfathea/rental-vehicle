package database

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"errors"
	"os"
	"log"
	"github.com/joho/godotenv"
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

	// config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable", host, user, pass, dbName)

	// gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	s := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, dbName)

	// dsn := "host=localhost user=vicky password=abcd1234 dbname=gocar sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{})

	if err != nil {
		return nil, errors.New("gagal konek db")
	  }

	// db, err := gormDb.DB()
	// if err != nil {
	// 	return nil, errors.New("gagal konek db")
	// }

	// db.SetConnMaxIdleTime(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxLifetime(time.Hour)

	return db, nil
}