package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandle)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error running server %v", err)
	}
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Hello")
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error while processing form r.ParseForm(): %v", err)
	}
	fmt.Println("Post request made")
	name := r.FormValue("name")
	add := r.FormValue("address")
	fmt.Fprintf(w, "Name got is %s with address as %s\n", name, add)
}
