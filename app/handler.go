package app

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Multiplexer struct {
	Template *template.Template
}

func (plexer *Multiplexer) IndexHandler(w http.ResponseWriter, req *http.Request) {

	// Check for cookies
	// If there is there is cookie
	c, err := req.Cookie("szn")
	if err != nil {
		// if no cookie give cookie
		uuid := uuid.Must(uuid.NewRandom())

		c = SendCookie(w, req, uuid.String())
	}

	if req.Method == http.MethodPost {
		// Parse request body as multipart form data with 32MB max memory
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// Get the uploaded file multipart data
		file, handler, err := req.FormFile("upload")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		defer file.Close()
		
		// Get the current working directory to create path to save the file
		cwd, err := os.Getwd()
		if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
		}
			
		// Hash the filename
		hashed := ExtractAndHash(handler.Filename)
		
		// Append the hashed name of the file to cookie
		nwC := fmt.Sprintf("%v|%v", c.Value, hashed)
		
		// Set new cook
		c = SendCookie(w, req, nwC)
		
		// Merge Directory to create file
		filePath := filepath.Join(cwd, "assets", "images", hashed)
		fmt.Println(filePath)
		
		// Create empty file with the file name
		fileUpload, err := os.Create(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer fileUpload.Close()

		// Copy the content of the upload content which is still in memory to the create file
		_, err = io.Copy(fileUpload, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(c.Value)

	}

	plexer.Template.ExecuteTemplate(w, "index.gohtml", c.Value)

}
