package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
	db_user  string = os.Getenv("MYSQL_USER")
	db_password  string = os.Getenv("MYSQL_PASSWORD")
	db_name  string = os.Getenv("MYSQL_DATABASE")
	db_host  string = os.Getenv("DATABASE_HOST")
	db_port  string = os.Getenv("DATABASE_PORT")

)


func Connect() {
	db_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	d, err := gorm.Open("mysql", db_url)
	if  err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}