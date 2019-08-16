package api

import (
	"fmt"
	"log"
	"net/http"
)

func apiIdex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "data Api")
}

func api(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("ok %v")
	w.Write([]byte(msg))
}

func handel(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Query OK")
	})
	http.HandleFunc("/apiindex", apiIndex)
	http.HandleFunc("/ok", api)

	log.Println("This port start at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
