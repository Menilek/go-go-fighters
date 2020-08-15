package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Fighter struct {
	Name string `json:"name"`
	Grappling string `json:"grappling"`
	Striking string `json:"striking"`
}

type Fighters []Fighter

func allFighters(w http.ResponseWriter, r *http.Request){
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
	http.HandleFunc("/", homePage)
	http.HandleFunc("/fighters", allFighters)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}