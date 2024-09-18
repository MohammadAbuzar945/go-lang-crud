# Book Management API with S3 Image Upload in GoLang

This is a Go-based CRUD API for managing books, which also allows uploading images to an AWS S3 bucket. The API follows a structured approach with different files for controllers, services, and configurations. The image files are uploaded with randomly generated names to avoid filename collisions.

## Features

- CRUD operations for managing books.
- Upload images to AWS S3 with a random filename.
- Uses PostgreSQL as the database.
- Structured project layout with controllers, services, and configuration files.
- Supports environment variables loaded from a `.env` file.

## Project Structure

## Requirements

- [Go](https://golang.org/dl/) (1.18 or higher)
- PostgreSQL
- AWS account with an S3 bucket
- AWS IAM user or role with access to the S3 bucket
- Environment variables set for AWS access

## Installation

1. Clone the repository:

```bash
git clone https://github.com/MohammadAbuzar945/go-lang-crud-with-s3bucket.git
cd go-lang-crud-with-s3bucket

## Set Up .Env
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
AWS_S3_BUCKET_NAME=your-bucket-name
AWS_REGION=eu-north-1
### Update Credentials in db.go
DB_HOST=localhost
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name



### Notes:
- Replace `Database configuration` and other placeholder values with your actual configuration.

