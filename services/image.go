package services

import (
	"bookManagement/config"
	s3 "bookManagement/config"
	"bookManagement/models"
	"mime/multipart"
	"time"
)

func CreateImage(bookID uint, file multipart.File, handler *multipart.FileHeader) (models.Image, error) {

	//read the file

	fileBytes := make([]byte, handler.Size)
	file.Read(fileBytes)

	//upload to s3
	imageURL, err := s3.UploadFilesToS3(handler.Filename, fileBytes)
	if err != nil {
		return models.Image{}, err
	}

	//save image to database

	image := models.Image{
		BookID:     bookID,
		ImageURL:   imageURL,
		UploadedAt: time.Now(),
	}
	err = config.DB.Create(&image).Error
	return image, err

}

func GetImagesByBookID(bookID uint) ([]models.Image, error) {
	var images []models.Image
	err := config.DB.Where("book_id = ?", bookID).Find(&images).Error
	return images, err
}
