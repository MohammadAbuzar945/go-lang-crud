package services

import (
	"bookManagement/config"
	"bookManagement/models"
)

func CreateBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

func GetBooks(books *[]models.Book) error {
	return config.DB.Find(books).Error
}
func GetBookById(book *models.Book, id string) error {
	return config.DB.First(book, id).Error
}

func UpdateBook(book *models.Book) error {
	return config.DB.Save(book).Error
}

func DeleteBook(book *models.Book) error {
	return config.DB.Delete(book).Error
}
