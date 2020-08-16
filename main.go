package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	Fighters []Fighter `json:"fighters"`
}

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

func apiFighters(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://top-fighters.herokuapp.com/api/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Fighters); i++ {
		json.NewEncoder(w).Encode(responseObject.Fighters[i].Name)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api", apiFighters).Methods("GET")
	myRouter.HandleFunc("/fighters", allFighters).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
