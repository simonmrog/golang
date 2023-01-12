package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type People struct {
	Id        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Fake database
var people []People

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPeopleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&People{})
}

func createPeople(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person People
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func deletePeople(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	// filling up database
	people = append(people, People{Id: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, People{Id: "2", FirstName: "Joe", LastName: "Ray"})

	// routes
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPeopleById).Methods("GET")
	router.HandleFunc("/people/{id}", createPeople).Methods("POST")
	router.HandleFunc("/people/{id}", deletePeople).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}
