package main

import (
	"fmt"
	"net/http"
)

// declaring Handler Funcs
func formHandler(w http.ResponseWriter, r *http.Request) {
	//we need to parse the form first and handle any possible errors
	if err := r.ParseForm(); err != nil {

		fmt.Fprintf(w, "There was an error parsing the form %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Your name is: %v \n", name)
	fmt.Fprintf(w, "Your address is: %v \n", address)
	fmt.Fprintf(w, "Thank you")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	//We don't want injections by Posting to the server
	if r.Method != "GET" {
		http.Error(w, "404", http.StatusNotFound)
	}

	fmt.Fprintf(w, "hello visitor")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	//declaring Handlefuncs
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//Starting Server
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", nil)

}
