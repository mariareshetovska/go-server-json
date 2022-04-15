package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json: "name"`
	Email string `json: "email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loadData("api/users.json"))
}

func loadData(filepath string) []User {
	file, _ := ioutil.ReadFile(filepath)
	var persons []User
	err := json.Unmarshal([]byte(file), &persons)

	if err != nil {
		fmt.Println("JSON decode error!")
	}
	return persons
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
