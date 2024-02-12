package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	bars "server1/api/v1"
)

// Server Main
func main() {

	log.Print("Server started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello!")
		if err != nil {
			log.Print(err)
		}
	})
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

// GET /foo
func foo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	fmt.Fprintf(w, "fooooooooooo")
}

// GET /bar
func bar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	json.NewEncoder(w).Encode(&bars.Bar{Name: "Bar Jam", NumOfTables: 3})
}
