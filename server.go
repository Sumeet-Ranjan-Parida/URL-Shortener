package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe("5050", router))

	router.HandleFunc("/shorten", shorturl).Methods("PUT")
}

func shorturl(w http.ResponseWriter, r *http.Request) {

}
