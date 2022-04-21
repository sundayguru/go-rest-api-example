package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/routes"
	"github.com/gorilla/mux"
)

func  homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}


func main()  {
	r := mux.NewRouter()
	routes.RegisterBankAccountRoutes(r)
	r.HandleFunc("/", homePage)
	fmt.Printf("Staring server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}