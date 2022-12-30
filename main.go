package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/templates/*"))
}

func main() {

	router := mux.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/", indexHandler)

	server.ListenAndServe()

}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	// Check for cookies
	// If there is there is cookie
	c, err := req.Cookie("szn")
	if err != nil {
		// if no cookie give cookie
		uuid := uuid.Must(uuid.NewRandom())

		c = sendCookie(w, req, uuid.String())
	}

	if req.Method == http.MethodPost {
		// Get the value from the form
		formValue := req.FormValue("upload")
		if formValue == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashed := extractAndHash(formValue)
		nwC := fmt.Sprintf("%v|%v", c.Value, hashed)

		c = sendCookie(w, req, nwC)
		
		fmt.Println(c.Value)

	}

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)

}

// ExtractAndHash extracts the name and extension of a file, hash the name then return the name as a hash sha1 value with the file extension.
func extractAndHash(input string) string {
	// extract the name and extension
	filename := strings.Split(input, ".")

	// hash the name
	hash := sha1.New()
	io.WriteString(hash, filename[0])

	hashed := fmt.Sprintf("%x", hash.Sum(nil))

	// append the extension back
	// append the value to the cookie
	return fmt.Sprintf("%v.%s", hashed, filename[1])
}


// Send Cookie sends a cookie to the http connection in context
func sendCookie(w http.ResponseWriter, r *http.Request, value string) *http.Cookie {
	c :=  &http.Cookie{
		Name:  "szn",
		Value: value,
	}

	http.SetCookie(w, c)

	return c
}
