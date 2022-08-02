package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imchukwu/book_management_system/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
