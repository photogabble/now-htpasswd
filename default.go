package handler

import (
	"fmt"
	"os"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Build protected file path from http request
	fp := filepath.Join("./protected", filepath.Clean(r.URL.Path))
	if fp == "protected" {
		fp = filepath.Join(fp, "index")
	}
	fp = fp + ".html"
	
	fmt.Fprintf(w, fp)
	
	return
	
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
	
	fmt.Fprintf(w, fp)
}
