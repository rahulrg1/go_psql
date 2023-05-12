package main

import (
	"fmt"
	"go_valid/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.GetAllData).Methods("GET")
	r.HandleFunc("/add", middleware.CreateForm).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("server started at 8000")
}
// all fields are required phonenumber should be int