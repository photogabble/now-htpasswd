package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("protected", filepath.Clean(r.URL.Path))
	
	if fp == "protected" {
		fp = fp+"index"
	}
	
	fp = fp + ".html"
	
	fmt.Fprintf(w, fp)
}
