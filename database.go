package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConnection connect to the databse
func DBConnection(schema string) (*gorm.DB, error) {
	// Create database connecttion
	db, err := gorm.Open("postgres", PsqlConn)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	// Show SQL script in terminal
	// db.LogMode(true)

	// Set table schema
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return schema + "." + tableName
	}

	return db, nil
}

// InitialMigration is the migration
func InitialMigration() {
	db, _ := DBConnection("demo")
	defer db.Close()

	db.AutoMigrate(&User{})
}
