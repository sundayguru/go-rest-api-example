package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
	DBUser  string = os.Getenv("MYSQL_USER")
	DBPassword  string = os.Getenv("MYSQL_PASSWORD")
	DBName  string = os.Getenv("MYSQL_DATABASE")
	DBHost  string = os.Getenv("DATABASE_HOST")
	DBPort  string = os.Getenv("DATABASE_PORT")
)


func Connect() {
	db_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	d, err := gorm.Open("mysql", db_url)
	if  err != nil {
		panic(err)
	}
	validations.RegisterCallbacks(d)
	db = d
}

func GetDB() *gorm.DB {
	return db
}