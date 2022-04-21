package main

import (
	"fmt"
	"log"
	"net/http"
)

func  homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}


func main()  {
	http.HandleFunc("/", homePage)

	fmt.Printf("starting server at port  8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}