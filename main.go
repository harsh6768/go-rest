package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book

//Book Struct Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct Model
type Author struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Inside get books method >>>")
	w.Header().Set("Content-Type", "application/json")
	//return data
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get id from request
	params := mux.Vars(r)

	for _, item := range books {

		if item.ID == params["id"] {
			//return data
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func createBook(res http.ResponseWriter, req *http.Request) {

	fmt.Printf("Request body >>>", req.body)

	res.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(req.body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)

	json.NewEncoder(res).Encode(book)

}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	//get id from request
	params := mux.Vars(req)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			return
		}
	}
	json.NewEncoder(res).Encode(books)

}

func main() {
	fmt.Printf("Hello harsh")
	// Init Router
	route := mux.NewRouter()

	//Mock Data -- @Todo - implment DB
	books = append(books, Book{
		ID:     "1",
		Isbn:   "1234",
		Title:  "Book1",
		Author: &Author{Firstname: "Harsh", LastName: "Chaurasiya"}})
	books = append(books, Book{
		ID:     "2",
		Isbn:   "1235",
		Title:  "Book2",
		Author: &Author{Firstname: "Shubham", LastName: "Singh"}})

	//routes
	route.HandleFunc("/api/getBooks", getBooks).Methods("GET")
	route.HandleFunc("/api/getBooks/{id}", getBookById).Methods("GET")
	route.HandleFunc("/api/createBook", createBook).Methods("POST")
	route.HandleFunc("/api/deleteBook/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", route))

}
