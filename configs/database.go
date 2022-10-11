package configs

import (
	"fmt"
	"golang-gorm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 5432
	user   string = "postgres"
	pass   string = "apayapasswordnya"
	dbname string = "hacktiv-go1"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Order{}) {
		db.AutoMigrate(&models.Order{})
	}

	if !db.HasTable(&models.Item{}) {
		db.AutoMigrate(&models.Item{})
	}

	return db
}
