package database

import (
	"log"

	"github.com/ricardomaricato/payment-routine/accounts-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect opens the database connection and returns it
func Connect() (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.Open(config.DataBaseConectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Error to connect to database")
		return nil, err
	}
	return DB, nil
}
