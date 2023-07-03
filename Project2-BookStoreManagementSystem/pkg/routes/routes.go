package routes

import (
	"github.com/ankit/Project2-BookManagementSystem/pkg/controller"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/Book", controller.CreateBook).Methods("POST")
	r.HandleFunc("/Book", controller.GetBooks).Methods("GET")
	r.HandleFunc("/Book/{bookId}", controller.GetBookById).Methods("GET")
	//r.HandleFunc("/Book/{bookId}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/Book/{bookId}", controller.DeleteBook).Methods("DELETE")
}
