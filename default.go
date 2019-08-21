package handler

import (
	"os"
	"io"
	"strconv"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
	
	// Open file for transmission
	oFile, err := os.Open(fp)
	defer oFile.Close()
	
	// If we can't open the file consider it not found
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	// Buffer for storing the file header
	oFileHeader := make([]byte, 512)
	// Copy 512 bytes into buffer
	oFile.Read(oFileHeader)
	// Analyse file header to obtain content type
	oFileContentType := http.DetectContentType(oFileHeader)
	// Get filesize as a string
	oFileSize := strconv.FormatInt(info.Size(), 10)
	
	// Set Response Headers
	w.Header().Set("Content-Type", oFileContentType)
	w.Header().Set("Content-Length", oFileSize)
	
	// Set Response Body from file content
	// Because we read the first 512 bytes into oFileHeader we need to reset
	// the file pointer offset back to zero
	oFile.Seek(0,0)
	
	io.Copy(w, oFile)
}
