package main

import (
	"log"
	"net/http"

	"shotlink/core"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://nexvul.com", http.StatusFound)
	})

	http.HandleFunc("/get", core.HandleScreenshot)
	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
