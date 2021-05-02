package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "uploading file\n")

	// 1. parse input , type multipart/form-data
	r.ParseMultipartForm(10)

	// 2. retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Upload file: %v+\n", handler.Filename)
	fmt.Printf("Size file: %v+\n", handler.Size)
	fmt.Printf("MIME file: %v+\n", handler.Header)

	// 3. write temporary file on your server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	// 4. return wheter or not this has been successful
	fmt.Fprintf(w, "successfully upload file\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":3000", nil)
}

func main() {
	fmt.Println("go upload file")
	setupRoutes()
}
