package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
)

type Url struct {
	ID       string `json:"id"`
	Shorturl string `json:"shorturl"`
	Longurl  string `json:"longurl"`
}

var urls []Url

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

func create(w http.ResponseWriter, r *http.Request) {
	var url Url
	_ = json.NewDecoder(r.Body).Decode(&url)

	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)

	now := time.Now()
	e, _ := h.Encode([]int{int(now.Unix())})

	//urls = append(urls, Url{ID: e, Shorturl: "http://localhost:8080/" + e, Longurl: url.Longurl})

	ID := e
	shorturl := "http://localhost:8080/" + e
	longurl := url.Longurl

	db := dbConnect()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO urls(id, shorturl, longurl) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(ID, shorturl, longurl)
}

func rootendpoint(w http.ResponseWriter, r *http.Request) {
	//URL Parameter
	urlparams := mux.Vars(r)

	id := urlparams["id"]

	db := dbConnect()
	defer db.Close()

	find, err := db.Query("SELECT longurl FROM urls WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	for find.Next() {
		var url Url
		err := find.Scan(&url.Longurl)
		if err != nil {
			panic(err.Error())
		}
		bigurl := "http://" + url.Longurl
		http.Redirect(w, r, bigurl, 301)
	}

}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", health)
	r.HandleFunc("/create", create).Methods("POST")
	r.HandleFunc("/{id}", rootendpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/urlshoten")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//fmt.Println("Connected")

	return db
}

func main() {
	fmt.Println("Server Started...")

	handleRequests()
}
