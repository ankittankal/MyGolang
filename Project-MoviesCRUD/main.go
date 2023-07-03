package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http" //allows us to create a server in golang
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"` //denote a struct tag, which provides additional metadata for the field. In this case, it indicates that when this struct is encoded or decoded to/from JSON, the field should be represented as id`.
	Isbn     string    `json : "isbn"`
	Title    string    `jsoon : "title"`
	Director *Director `json : director`
}

type Director struct {
	Firstname string `json: "firstname`
	Lastname  string `json : lastname`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	//fields := strings.Split(r.URL.Path, "/")
	//movieID := fields[2]
	params := mux.Vars(r)

	for index, item := range movies {

		//if item.ID == movieID {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fields := strings.Split(r.URL.Path, "/")
	movieID := fields[2]

	for _, item := range movies {

		if item.ID == movieID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn((10000000)))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("content-type", "application/json")

	//extract id from url path
	fields := strings.Split(r.URL.Path, "/")
	movieID := fields[2]

	//loop over the movies
	//delete the movie with ID that you have sent
	//add a new movie - the movie which we send in the body of postman

	for index, item := range movies {

		if item.ID == movieID {
			movies = append(movies[:index], movies[index+1:]...)
			var newMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = movieID
			movies = append(movies, newMovie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie one", Director: &Director{Firstname: "john", Lastname: "Deo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "56789", Title: "Movie two", Director: &Director{Firstname: "steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server st port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
