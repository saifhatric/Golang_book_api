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
type Author struct{
	Name 		string	`json:"Name"`
	Lastname 	string	`json:"lastname"`
	Age 		int		`json:"age"`
}
type Books struct{
	ID			string `json:"id"`
	Title 		string	`json:"title"`
	Quantity 	int		`json:"quantity"`
	Author		*Author
	
}

var books = []Books{
	{ID: "1",Title: "Horney",Quantity: 200,Author: &Author{Name:"saif",Lastname:"benzamit",Age:19}},
	{ID: "10",Title: "Lust",Quantity: 400,Author: &Author{Name:"said",Lastname:"benzamit",Age:20}},
	{ID: "101",Title: "Cool",Quantity: 404,Author: &Author{Name:"rachid",Lastname:"benzamit",Age:21}},
}

func main(){
	fmt.Println("Server started!");
	router:= mux.NewRouter()
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/book",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080",router))
}

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	err:=json.NewEncoder(w).Encode(books)
	ErrorHandler(err)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	prams:=mux.Vars(r) // Get any prams
	for _,book := range books{
		if(prams["id"] == book.ID){
			err:=json.NewEncoder(w).Encode(book)
			ErrorHandler(err)
			return
		}
	}
	json.NewEncoder(w).Encode("Sorry this book doesn't exist!")
}
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var newBook Books
	json.NewDecoder(r.Body).Decode(&newBook)
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)

}
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	var upBook Books
	prams:= mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&upBook)
	upBook.ID = strconv.Itoa(rand.Intn(100000))
	for index,book:= range books{
		if(prams["id"]==book.ID){
			books= append(books[:index],books[index+1:]...)
			books = append(books, upBook)
			json.NewEncoder(w).Encode(upBook)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	prams:=mux.Vars(r)
	for index,book:= range books{
		if(prams["id"]==book.ID){
			books = append(books[:index],books[index+1:]...)
			break 
		}
	}
	json.NewEncoder(w).Encode(&books)

}

func ErrorHandler(err error){
	if err!=nil{
		log.Fatal(err)
	}
}