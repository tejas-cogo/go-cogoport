package models

import (
	"gorm.io/gorm"
	// "github.com/tejas-cogo/go-cogoport/config"
)

var db *gorm.DB

//Model is sample of common table structure

func init() {

	// config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model
	db.Migrator().AutoMigrate(
		&GroupMember{},
	)
}

// GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
