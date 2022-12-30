package main

import (
	"html/template"

	"photo.dev/app"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/templates/*"))
}

func main() {
	Application := app.App{}
	Application.Init(tpl)
	Application.Run()

	// router := mux.NewRouter()

	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// router.HandleFunc("/", indexHandler)

	// server.ListenAndServe()

}

// func indexHandler(w http.ResponseWriter, req *http.Request) {

// 	// Check for cookies
// 	// If there is there is cookie
// 	c, err := req.Cookie("szn")
// 	if err != nil {
// 		// if no cookie give cookie
// 		uuid := uuid.Must(uuid.NewRandom())

// 		c = sendCookie(w, req, uuid.String())
// 	}

// 	if req.Method == http.MethodPost {
// 		// Parse request body as multipart form data with 32MB max memory
// 		err := req.ParseMultipartForm(32 << 20)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		// Get the uploaded file multipart data
// 		file, handler, err := req.FormFile("upload")
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		defer file.Close()

// 		  // Get the current working directory to create path to save the file
// 		cwd, err := os.Getwd()
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		// Hash the filename
// 		hashed := extractAndHash(handler.Filename)

// 		// Append the hashed name of the file to cookie
// 		nwC := fmt.Sprintf("%v|%v", c.Value, hashed)

// 		// Set new cook
// 		c = sendCookie(w, req, nwC)
		

// 		// Merge Directory to create file
// 		filePath := filepath.Join(cwd, "assets", "images", hashed)

// 		// Create empty file with the file name
// 		fileUpload, err := os.Create(filePath)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		defer fileUpload.Close()

// 		  // Copy the content of the upload content which is still in memory to the create file
// 		_, err = io.Copy(fileUpload, file)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
		
// 		fmt.Println(c.Value)

// 	}

// 	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)

// }

// // ExtractAndHash extracts the name and extension of a file, hash the name then return the name as a hash sha1 value with the file extension.
// func extractAndHash(input string) string {
// 	// extract the name and extension
// 	filename := strings.Split(input, ".")

// 	// hash the name
// 	hash := sha1.New()
// 	io.WriteString(hash, filename[0])

// 	hashed := fmt.Sprintf("%x", hash.Sum(nil))

// 	// append the extension back
// 	// append the value to the cookie
// 	return fmt.Sprintf("%v.%s", hashed, filename[1])
// }


// // Send Cookie sends a cookie to the http connection in context
// func sendCookie(w http.ResponseWriter, r *http.Request, value string) *http.Cookie {
// 	c :=  &http.Cookie{
// 		Name:  "szn",
// 		Value: value,
// 	}

// 	http.SetCookie(w, c)

// 	return c
// }
