package controllers

import (
	"encoding/json"
	"fmt"
	Img_model "golang_upload_file/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Respone struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// https://tutorialedge.net/golang/go-file-upload-tutorial/
	if r.Method == "POST" {

		fmt.Println("File Upload Endpoint Hit")
		db, err := Img_model.ConnectDB()
		fmt.Println("error database xxxxxxxxxx")
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			var rs = Respone{
				Status:  http.StatusForbidden,
				Message: err.Error(),
			}
			fmt.Println(rs)
			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusForbidden)
			w.Write(respons_data)
			return
		}
		defer db.Close()

		token := r.FormValue("token")
		if token != os.Getenv("TOKEN") {
			var rs = Respone{
				Status:  http.StatusForbidden,
				Message: "Token not matches",
			}
			fmt.Println(rs)
			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusForbidden)
			w.Write(respons_data)
			return
		}
		// FormFile returns the first file for the given key `file`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("file")
		if err != nil {
			var rs = Respone{
				Status:  http.StatusForbidden,
				Message: "Error Retrieving the File",
			}
			fmt.Println(rs)
			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusForbidden)
			w.Write(respons_data)
			return
		}
		defer file.Close()
		extension := strings.Split(handler.Filename, ".")[len(strings.Split(handler.Filename, "."))-1]
		ex_lower := strings.ToLower(extension)

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header["Content-Type"])
		fmt.Printf("extension : %+v\n", ex_lower)

		//transform in MB
		var bytes = handler.Size
		var kilobytes = (float64)(bytes / 1024)
		var megabytes = (float64)(kilobytes / 1024)
		if megabytes > 10 {
			var rs = Respone{
				Status:  http.StatusForbidden,
				Message: "The uploaded file is too big.",
			}
			fmt.Println(rs)
			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusForbidden)
			w.Write(respons_data)
			return
		}
		fmt.Println("cctyiiii")
		fmt.Println(handler.Header["Content-Type"][0])
		// if handler.Header["Content-Type"][0] == "image/png" || handler.Header["Content-Type"][0] == "image/jpeg" {
		if ex_lower == "png" || ex_lower == "jpg" || ex_lower == "jpeg" {
			fmt.Println("Yes")

			// Create a temporary file within our temp-images directory that follows
			// a particular naming pattern
			tempFile, err := ioutil.TempFile("temp-images/", "upload-*."+ex_lower)

			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()

			// read all of the contents of our uploaded file into a
			// byte array
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}

			imagePath := tempFile.Name()
			contentType := handler.Header["Content-Type"]
			imagePathNew := strings.Replace(imagePath, "\\", "/", 1)
			imageSize := strconv.Itoa(int(handler.Size))
			err = Img_model.CreateImageData(db, contentType, imageSize, imagePathNew)
			if err != nil {
				fmt.Println(err)
				var rs = Respone{
					Status:  http.StatusForbidden,
					Message: err.Error(),
				}
				fmt.Println(rs)
				var respons_data, _ = json.Marshal(rs)
				w.WriteHeader(http.StatusForbidden)
				w.Write(respons_data)
				return
			}
			// write this byte array to our temporary file
			tempFile.Write(fileBytes)
			var rs = Respone{
				Status:  http.StatusOK,
				Message: "Successfully Uploaded File.",
			}

			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusOK)
			w.Write(respons_data)
			return
		} else {
			var rs = Respone{
				Status:  http.StatusForbidden,
				Message: "Not support file " + ex_lower,
			}
			var respons_data, _ = json.Marshal(rs)
			w.WriteHeader(http.StatusForbidden)
			w.Write(respons_data)
			return
		}

	} else {

		var rs = Respone{
			Status:  http.StatusMethodNotAllowed,
			Message: "Method Not Allowed",
		}
		var respons_data, _ = json.Marshal(rs)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(respons_data)
	}
}
