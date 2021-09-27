package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ImageDataPost struct {
	Content_type string `json:"content_type"`
	Size         string `json:"size"`
	Image_path   string `json:"image_path"`
}

func init() {
	// https://zetcode.com/golang/env/
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func ConnectDB() (*sql.DB, error) {

	sqlHost := os.Getenv("DB_HOST")
	sqlPort := os.Getenv("DB_PORT")
	sqlType := os.Getenv("DB_DRIVER")
	sqlUsn := os.Getenv("DB_USER")
	sqlPass := os.Getenv("DB_PASS")
	sqlDatabase := os.Getenv("DB_DATABASE")

	db, err := sql.Open(sqlType, sqlUsn+":"+sqlPass+"@tcp("+sqlHost+":"+sqlPort+")/"+sqlDatabase)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateImageData(db *sql.DB, Content_Type []string, Size string, Image_path string) error {
	_, err := db.Query("INSERT INTO image_tb (content_type,size,image_path) VALUES ('" + Content_Type[0] + "','" + Size + "','" + Image_path + "')")
	if err != nil {
		return err
	}
	return nil
}
