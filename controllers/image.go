package controllers

import (
	"bookManagement/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.Atoi(params["book_id"])

	fmt.Println("BookId", bookID)

	if err != nil {
		http.Error(w, "Invalid Book Id", http.StatusBadRequest)
		return
	}

	// fmt.Println("Form", r.FormFile("file"))
	file, handler, err := r.FormFile("file")
	fmt.Println("File", file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	image, err := services.CreateImage(uint(bookID), file, handler)
	fmt.Println("Image", image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&image)

}

// GetImagesHandler retrieves images for a specific book
func GetImagesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse book ID from URL
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["book_id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Call service to get images
	images, err := services.GetImagesByBookID(uint(bookID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&images)
}
