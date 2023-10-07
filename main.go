package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/products", ProductsPage)
	http.HandleFunc("/style.css", StyleCSS)
	http.HandleFunc("/get_products.js", GetProductScript)
	http.HandleFunc("/api/product", ProductService)
	
	

	log.Fatal(http.ListenAndServe(":8080", nil))
}
