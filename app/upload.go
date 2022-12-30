package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)


func uploadFile(w http.ResponseWriter, req *http.Request, c *http.Cookie) (*http.Cookie, error) {
	// Parse request body as multipart form data with 32MB max memory
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, fmt.Errorf("Error")
	}
	
	// Get the uploaded file multipart data
	file, handler, err := req.FormFile("upload")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, fmt.Errorf("Error")
	}
	
	defer file.Close()
	
	// Get the current working directory to create path to save the file
	cwd, err := os.Getwd()
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return nil, fmt.Errorf("Error")
	}
		
	// Hash the filename
	hashed := ExtractAndHash(handler.Filename)
	
	// Append the hashed name of the file to cookie
	nwC := fmt.Sprintf("%v|%v", c.Value, hashed)
	
	// Set new cook
	c = SendCookie(w, req, nwC)
	
	// Merge Directory to create file
	filePath := filepath.Join(cwd, "assets", "images", hashed)
	
	// Create empty file with the file name
	fileUpload, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, fmt.Errorf("Error")
	}

	defer fileUpload.Close()

	// Copy the content of the upload content which is still in memory to the create file
	_, err = io.Copy(fileUpload, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, fmt.Errorf("Error")
	}

	return c, nil

}
