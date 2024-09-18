package controllers

import (
	"bookManagement/models"
	"bookManagement/services"
	"encoding/json"
	"log"
	"net/http"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if err := services.CreateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	log.Println("Book created successfully")
	json.NewEncoder(w).Encode(book)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := services.GetBooks(&books); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(books)

}

func GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	id := r.URL.Query().Get("id")

	println(id)
	if err := services.GetBookById(&book, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	// id := r.URL.Query().Get("id")
	json.NewDecoder(r.Body).Decode(&book)
	if err := services.UpdateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(book)

}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	if err := services.DeleteBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(book)
}
