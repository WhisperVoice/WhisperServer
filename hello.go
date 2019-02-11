package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/whispers", retrieveWhispersHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST");
	params := mux.Vars(r)
	fmt.Print(params)
	fmt.Println("_______________")
	fmt.Print(r.Body)
}

// TODO: Take sign up registration informationn and register to cosmosdb table https://docs.microsoft.com/en-us/rest/api/storageservices/table-service-rest-api
func signupHandler(w http.ResponseWriter, r *http.Request) {
}

// TODO: Retrieve voice clips for user post endpoint https://docs.microsoft.com/en-us/rest/api/storageservices/table-service-rest-api
func retrieveWhispersHandler(w http.ResponseWriter, r *http.Request) {
}

// TODO: Retrieve voice clips for user mainn function https://docs.microsoft.com/en-us/rest/api/storageservices/table-service-rest-api
func retrieveWhispers(w http.ResponseWriter, r *http.Request) {
}
