package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home !")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Todo !")
	t, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	fmt.Println("server is starting on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
