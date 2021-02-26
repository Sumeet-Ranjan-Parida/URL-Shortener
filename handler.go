package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
)

type Url struct {
	ID       string `json:"id"`
	Shorturl string `json:"shorturl"`
	Longurl  string `json:"longurl"`
}

var urls []Url

func getUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)

	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)

	now := time.Now()
	e, _ := h.Encode([]int{int(now.Unix())})

	//Sample Data
	urls = append(urls, Url{ID: e, Shorturl: "localhost:8080/" + e, Longurl: "long"})
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/urls", getUrls).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Server Started...")

	handleRequests()

	db := dbConnect()
}
