package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DBConnection connect to the databse
func DBConnection(schema string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", PsqlConn)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	db.LogMode(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return schema + "." + tableName
	}

	return db, nil
}
