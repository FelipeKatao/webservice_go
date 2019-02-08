package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// função principal
func main() {
    router := mux.NewRouter()
    
    router := mux.NewRouter()
    router.HandleFunc("/contato", GetPeople).Methods("GET")
    router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
    
    log.Fatal(http.ListenAndServe(":8000", router))
}
