package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLStore struct {
	db *gorm.DB
}

func New(dsn string) MySQLStore {
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failed to connect to database!")
	}
	if aErr := database.AutoMigrate(&User{}, &Ticket{}, &Order{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	return MySQLStore{db: database}
}
