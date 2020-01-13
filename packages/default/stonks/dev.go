
//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
)

// This main function is for testing locally
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query()