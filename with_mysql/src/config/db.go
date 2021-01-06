package config

/** This package contains app wide re-usable db connection **/

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB connection
func DB() *gorm.DB {

	db, err := gorm.Open("mysql", Config("MySQL_DB_USER")+":"+Config("MySQL_DB_PASSWORD")+"@("+Config("MySQL_DB_HOST")+":"+Config("MySQL_DB_PORT")+")/"+Config("MySQL_DB_NAME")+"?charset="+Config("MySQL_DB_CHARSET")+"&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println("DB Connection Failure: %", err.Error())
		panic(err)
	}

	//if connection is successful without any errors
	fmt.Println("Connected to MySQL!") //comment this out in production

	//if connection is okay we return the db instance
	return db
}
