package handler

import (
	"os"
	"net/http"
	"path/filepath"
	"strings"
	"encoding/base64"
	"github.com/tg123/go-htpasswd"
)

var auth *htpasswd.File

func authenticate(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}
	
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}
	
	ok := auth.Match(pair[0], pair[1])
	
	if ok {
		w.Header().Set("X-Authenticated-Username", pair[0])
	}
	
	return ok
}

func Handler(w http.ResponseWriter, r *http.Request) {	
	// Load .htpasswd file
	auth = htpasswd.New("./.htpasswd", htpasswd.DefaultSystems, nil)
	
	// Build protected file path from http request
	fp := filepath.Join("protected", filepath.Clean(r.URL.Path))
	if fp == "protected" {
		fp = filepath.Join(fp, "index")
	}
	fp = "./" + fp + ".html"
		
	// Return a 404 if the protected file does not exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}
	
	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
	
	// Check if user is authenticated
	if authenticate(w,r,username,password) {
		http.ServeFile(w,r,fp)
		return
	}
	
	w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
        w.WriteHeader(401)
        w.Write([]byte("401 Unauthorized\n"))
}
