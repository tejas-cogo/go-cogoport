package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("database_name")
	password := os.Getenv("database_password")
	dbName := os.Getenv("database_username")
	dbHost := os.Getenv("database_host")
	dbPort := os.Getenv("database_port")

	postgres := postgres.Config{}
	fmt.Println(postgres)

	conn, err := gorm.Open("postgres", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")
	if err != nil {
		fmt.Print(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
