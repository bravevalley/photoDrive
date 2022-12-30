package app

import (
	"html/template"
	"net/http"

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

	// If the client uploads a file
	if req.Method == http.MethodPost {
		c, err = uploadFile(w, req, c)
		if err != nil {
			return
		}
	}

	// Get the cookie and split the values from the cookie
	uploads := SplitAndRetrieve(c.Value)

	plexer.Template.ExecuteTemplate(w, "index.gohtml", uploads)

}
