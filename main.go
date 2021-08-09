package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// A global variable that is incremented everytime a book is added.
// Used for providing a unique ID to each book
var count int

type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// A slice that will contain the books
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = count
	count++

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	vars := mux.Vars(r)

	var tmpBook Book
	for index, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			_ = json.NewDecoder(r.Body).Decode(&tmpBook)
			tmpBook.ID = index
			books[index] = tmpBook
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	vars := mux.Vars(r)
	var tmpBook Book
	for index, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			tmpBook = books[index]
			books = append(books[:index], books[index+1:]...)
			_ = json.NewEncoder(w).Encode(tmpBook)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {

	author1 := Author{FirstName: "Stanislaw", LastName: "Lem"}
	author2 := Author{FirstName: "Arthur", LastName: "Clarke"}

	count = 0

	books = append(books, Book{ID: count, Isbn: "123456", Title: "Solaris", Author: &author1})
	count++

	books = append(books, Book{ID: count, Isbn: "456280", Title: "2001: A Space Odyssey", Author: &author2})
	count++

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", addBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
