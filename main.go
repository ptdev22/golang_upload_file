package main

import (
	"fmt"
	"golang_upload_file/controllers"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
)

type PageData struct {
	PageTitle string
	Token     string
}

func init() {
	// https://zetcode.com/golang/env/
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	// https://gowebexamples.com/templates/
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Test upload file",
			Token:     os.Getenv("TOKEN"),
		}
		tmpl.Execute(w, data)
	})
	openBrowser("http://127.0.0.1:8080/")
	http.HandleFunc("/Upload", controllers.Upload)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
