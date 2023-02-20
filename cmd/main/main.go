package main

import (
	"fmt"
	"log"
	"myGolangProjects/golang-bookstore/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running at port 9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
