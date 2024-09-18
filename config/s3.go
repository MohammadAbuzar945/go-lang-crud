package config

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var S3Client *s3.Client
var RegionName = "eu-north-1"

// Generate a random string for the filename
func generateRandomString(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func InitilizeS3() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables after loading .env
	BucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	AccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	SecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if BucketName == "" || AccessKey == "" || SecretKey == "" {
		log.Fatal("Missing required AWS environment variables")
	}

	// Initialize S3 client with credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(RegionName),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			AccessKey,
			SecretKey,
			"", // session token, leave empty if not using one
		)),
	)
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	S3Client = s3.NewFromConfig(cfg)

	fmt.Println("S3 initialized with Bucket:", BucketName)
	if BucketName == "" {
		log.Fatal("Bucket name is required")
	}
}

func UploadFilesToS3(filename string, fileBytes []byte) (string, error) {
	// Retrieve bucket name from environment variables inside the function
	BucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	randomFileName, err := generateRandomString(16) // Generates a 16-byte random string
	if err != nil {
		return "", fmt.Errorf("failed to generate random file name: %v", err)
	}

	// Add file extension to random file name if necessary
	filename = randomFileName

	input := &s3.PutObjectInput{
		Bucket: &BucketName,
		Key:    &filename,
		Body:   bytes.NewReader(fileBytes),
	}

	_, err = S3Client.PutObject(context.TODO(), input)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", BucketName, RegionName, filename)
	return url, nil
}
