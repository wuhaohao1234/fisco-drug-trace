package main

/*
#include <signal.h>
*/
import "C"

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

// UploadHandler handles the file upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the multipart form
    err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusInternalServerError)
        return
    }

    // Get the file from the form
    file, handler, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Unable to get file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Create a new file
    f, err := os.OpenFile(filepath.Join("uploads", handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        http.Error(w, "Unable to create file", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    // Copy the file to the destination
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Unable to copy file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File uploaded successfully")
}

func main() {
    // Create the uploads directory if it doesn't exist
    if _, err := os.Stat("uploads"); os.IsNotExist(err) {
        os.Mkdir("uploads", 0755)
    }

    // Serve static files from the uploads directory
    http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

    // Handle the file upload
    http.HandleFunc("/upload", UploadHandler)

    // Start the server
    fmt.Println("Server is running on :3000")
    http.ListenAndServe(":3000", nil)
}
