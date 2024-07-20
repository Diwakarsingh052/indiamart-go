package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 * 1024 * 1024) // 10MB
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, header, err := r.FormFile("file-data")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a file on local file system to write the uploaded file content
	dstFile, err := os.Create("form/" + header.Filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create destination file", http.StatusInternalServerError)
		return
	}
	defer dstFile.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dstFile, file); err != nil {
		http.Error(w, "Failed to copy file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File uploaded successfully : ", header.Filename)
}

func main() {
	fmt.Println("Started server")
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

/*
1. Open Postman
2. Click on "New" > "Request"
4. In the opened request tab, select "POST" as the HTTP method from the dropdown menu next to the address bar
5. In the address bar, enter the URL of the endpoint where the file should be sent.
6. Below the address bar, click on the "Body" tab
7. Select "form-data"
8. In the "Key" field, type "file". In the dropdown on the right (which may currently say "Text"), select "File"
9. In the "Value" field, click "Select Files" and choose the file you want to upload from your file system
10. Click "Send" to send the file
*/
