package app

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ExtractAndHash extracts the name and extension of a file, hash the name then return the name as a hash sha1 value with the file extension.
func ExtractAndHash(input string) string {
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
func SendCookie(w http.ResponseWriter, r *http.Request, value string) *http.Cookie {
	c := &http.Cookie{
		Name:  "szn",
		Value: value,
	}

	http.SetCookie(w, c)

	return c
}
