package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server in :8080")

	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "HELLO!")
}


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s\n", name)

	fmt.Fprintf(w, "Address : %s\n", address)
}
