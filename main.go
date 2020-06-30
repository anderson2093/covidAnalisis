package main

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", Index).Methods("GET")
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
