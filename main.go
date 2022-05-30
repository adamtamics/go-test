package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(getenv("PORT", ":8081"), nil))

}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
			return fallback
	}
	return value
}