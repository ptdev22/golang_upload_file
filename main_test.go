package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var token_jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6IlRoaXMgaXMgdG9rZW4ifQ.CHrLEv2-DYi0seWoudiJMVI82tSB4hxX_f1kD_aFGrE"

func callUploadApi(token, path string) int {
	url := "http://127.0.0.1:8080/Upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	file, _ := os.Open(path)
	defer file.Close()
	part2, _ := writer.CreateFormFile("file", filepath.Base(path))
	_, errFile2 := io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return res.StatusCode
}

func Test_token_wrong(t *testing.T) {

	status := callUploadApi("xxx", "./image_test/img_1mb.jpg")
	// fmt.Println(string(body))
	if status != 403 {
		t.Errorf("Status code should be %v but got %v", 403, status)
	}
}
func Test_Content_type_wrong(t *testing.T) {

	status := callUploadApi(token_jwt, "./image_test/gif_file.gif")
	// fmt.Println(string(body))
	if status != 403 {
		t.Errorf("Status code should be %v but got %v", 403, status)
	}
}
func Test_over_size(t *testing.T) {

	status := callUploadApi(token_jwt, "./image_test/img_10mb.jpg")
	// fmt.Println(string(body))
	if status != 403 {
		t.Errorf("Status code should be %v but got %v", 403, status)
	}
}

func Test_upload_success(t *testing.T) {

	status := callUploadApi(token_jwt, "./image_test/img_1mb.jpg")
	// fmt.Println(string(body))
	if status != 200 {
		t.Errorf("Status code should be %v but got %v", 200, status)
	}
}
