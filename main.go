package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Fighter struct {
	Name      string `json:"name"`
	Grappling string `json:"grappling"`
	Striking  string `json:"striking"`
}

type Fighters []Fighter

func allFighters(w http.ResponseWriter, r *http.Request) {
	fighters := Fighters{
		Fighter{Name: "Menilek", Grappling: "Jiu Jitsu", Striking: "Muay Thai"},
	}

	fmt.Println("Top Fighters")
	json.NewEncoder(w).Encode(fighters)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/fighters", allFighters).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
