package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/jhidalgo3/containerized-golang-and-vuejs/model"
)

// UsersIndex returns index page with users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

// User Struct is the model for the app
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

// GetUsers returns json payload with users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.Users
	var payload []byte
	var err error

	//users = model.GetUsers()

	context := model.GetContext()
	c := context.DBCollection()
	c.Find(nil).All(&users)

	payload, err = json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// AddUser add a new user to database
func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.UserModel
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	userModel := model.UserModel{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Gender:    u.Gender,
	}
	context := model.GetContext()
	c := context.DBCollection()
	err2 := c.Insert(&userModel)
	if err2 != nil {
		log.Fatal("error saving user")
		return
	}

	code := struct {
		StatusCode int
	}{
		200,
	}
	payload, err3 := json.Marshal(code)
	if err3 != nil {
		log.Fatal("Something went wrong marshalling")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

// GetUserByID returns a user from mongo
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var id map[string]string
	err := json.NewDecoder(r.Body).Decode(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var result model.UserModel
	context := model.GetContext()
	c := context.DBCollection()
	user := c.Find(bson.M{"id": id["id"]}).All(&result)
	payload, err2 := json.Marshal(user)
	if err2 != nil {
		log.Fatal("Something went wrong marshalling")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

// DeleteUserByID deletes a user by id
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		ID int `json:"id"`
	}
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	fmt.Println(payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	context := model.GetContext()
	c := context.DBCollection()
	err2 := c.Remove(bson.M{"id": payload.ID})
	if err2 != nil {
		http.Error(w, err2.Error(), 500)
		return
	}

	code := struct {
		StatusCode int
	}{
		204,
	}
	payload2, err4 := json.Marshal(code)
	if err4 != nil {
		http.Error(w, err4.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload2)
}
