package routes

import (
	"bookManagement/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", controllers.CreateBookHandler).Methods(http.MethodPost)
	r.HandleFunc("/books", controllers.GetBooksHandler).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", controllers.GetBookByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/books", controllers.UpdateBookHandler).Methods(http.MethodPut)
	r.HandleFunc("/books", controllers.DeleteBookHandler).Methods(http.MethodDelete)

	r.HandleFunc("/books/{book_id}/images", controllers.UploadImageHandler).Methods(http.MethodPost)
	r.HandleFunc("/books/{book_id}/images", controllers.GetImagesHandler).Methods(http.MethodGet)

	return r
}
