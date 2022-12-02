package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 status not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 status not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World!!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() : %v\n", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	srv := &http.Server{
		Addr: "localhost:5000",
	}
	fmt.Printf("Starting server at port 5000")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
