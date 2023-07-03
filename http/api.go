package main

import (
	"fmt"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// // Planet is a planet type
// type Planet struct {
// 	Name       string `json:"name"`
// 	Population string `json:"population"`
// 	Terrain    string `json:"terrain"`
// }

// // Person is a person type
// type Person struct {
// 	Name         string `json:"name"`
// 	HomeworldURL string `json:"homeworld"`
// 	Homeworld    Planet
// }

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}
	fmt.Println(res)

}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
