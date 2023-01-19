package config

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	dbc *gorm.DB
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbPort, username, password, dbName)

	conn, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func ConnectCogoport() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("cogoport_database_username")
	password := os.Getenv("cogoport_database_password")
	dbName := os.Getenv("cogoport_database_name")
	dbHost := os.Getenv("cogoport_database_host")
	dbPort := os.Getenv("cogoport_database_port")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbPort, username, password, dbName)

	conn, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}
	dbc = conn
}

func GetDB() *gorm.DB {
	return db
}

func GetCDB() *gorm.DB {
	return dbc
}
