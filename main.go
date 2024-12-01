package main

import (
	"os"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	port := os.Getenv("PORT")
 
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
