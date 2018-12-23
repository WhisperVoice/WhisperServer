package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Print(params)
	fmt.Println("_______________")
	fmt.Print(r.Body)
}
