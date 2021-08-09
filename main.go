package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A global variable that is incremented everytime a book is added.
// Used for providing a unique ID to each book
var count int

type Book struct {
	ID     int     `json:"id"`
	Isdn   string  `json:"isdn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// A slice that will contain the books
var books []Book

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Run Successfully!</h1>")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", router))
}
