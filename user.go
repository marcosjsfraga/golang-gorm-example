package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// User structure
type User struct {
	gorm.Model
	Name  string
	Email string
	// Phone string
}

// CreateUser create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var user User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		fmt.Fprintf(w, "BadRequest")
		return
	}

	db, _ := DBConnection(domainParts[0])
	defer db.Close()

	db.Create(&User{
		Name:  user.Name,
		Email: user.Email,
	})

	fmt.Fprintf(w, "User successfully created")
}

// AllUsers returns all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	db, _ := DBConnection(domainParts[0])
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

// UpdateUser delete a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	params := mux.Vars(r)
	id := params["id"]

	fmt.Println(id)

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var requestUser User
	if err = json.Unmarshal(requestBody, &requestUser); err != nil {
		fmt.Fprintf(w, "BadRequest")
		return
	}

	db, _ := DBConnection(domainParts[0])
	defer db.Close()

	var user User
	db.Where("id=?", id).Find(&user)

	fmt.Println(&user)

	user.Name = requestUser.Name
	user.Email = requestUser.Email
	db.Save(&user)

	fmt.Fprintf(w, "User succesfully updated")
}

// DeleteUser delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	params := mux.Vars(r)
	id := params["id"]

	db, _ := DBConnection(domainParts[0])
	defer db.Close()

	var user User
	db.Where("id=?", id).Find(&user)

	db.Delete(&user)

	fmt.Fprintf(w, "User succesfully deleted")
}
