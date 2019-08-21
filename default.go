package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("protected", filepath.Clean(r.URL.Path))
	fmt.Fprintf(w, fp)
}
