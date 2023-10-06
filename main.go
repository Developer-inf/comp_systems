package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/api/product", ProductService)
	
	

	log.Fatal(http.ListenAndServe(":8080", nil))
}
