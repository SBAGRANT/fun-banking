package persistence

import (
	"log"
	"os"

	"github.com/bytebury/fun-banking/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database")
	}

	DB = db
}

func RunMigrations() {
	DB.AutoMigrate(&domain.User{})

	log.Println("🟢 successfully ran the migrations")
}