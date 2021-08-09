package config

import (
	"ai4am.com/go-restapi/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
//
// parameters:
// 		config DatabaseConfiguration
func SetupDB(config DatabaseConfiguration) {
	var db = DB

	driver := config.Driver
	database := config.Dbname
	username := config.Username
	password := config.Password
	host := config.Host
	port := config.Port

	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open("sqlite3", "./"+database+".db")
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "postgres" { // POSTGRES
		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			fmt.Println("db err: ", err)
		}
	} else if driver == "mysql" { // MYSQL
		db, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	// Change this to true if you want to see SQL queries
	db.LogMode(false)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Task{})
	DB.AutoMigrate(&models.Item{})
}

func GetDB() *gorm.DB {
	return DB
}