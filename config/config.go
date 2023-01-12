package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	_"github.com/jinzhu/gorm/dialects/postgres"
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

	username := os.Getenv("database_username")
	password := os.Getenv("database_password")
	dbName := os.Getenv("database_name")
	dbHost := os.Getenv("database_host")
	dbPort := os.Getenv("database_port")

	// postgres_config := postgres.Config{}
	// fmt.Println(postgres)

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbPort, username, password, dbName)

	// dsn := "host=} user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	conn, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})

	// conn, err := gorm.Open("postgres", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")
	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
