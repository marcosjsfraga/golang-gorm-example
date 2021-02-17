package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User structure
type User struct {
	gorm.Model
	Name  string
	Email string
}

// InitialMigration is the migration
func InitialMigration() {
	db, _ := DBConnection("demo")
	db.AutoMigrate(&User{})
}

// CreateUser create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var user User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		fmt.Fprintf(w, "BadRequest")
		return
	}

	db, _ := DBConnection("demo")
	defer db.Close()

	db.Create(&User{
		Name:  user.Name,
		Email: user.Email,
	})

	fmt.Fprintf(w, "User successfully created")
}

// AllUsers returns all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, _ := DBConnection("demo")

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

// UpdateUser delete a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := DBConnection("demo")
	defer db.Close()

	var user User
	db.Where("id=?", id).Find(&user)
	db.Update(&user)

	fmt.Fprintf(w, "User succesfully updated")
}

// DeleteUser delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, _ := DBConnection("demo")
	defer db.Close()

	var user User
	db.Where("id=?", id).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User succesfully deleted")
}
