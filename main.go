package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	main_page           = "/"
	products_page       = "/products"
	style               = "/style.css"
	get_products_script = "/get_products.js"
	api_product         = "/api/product"
)

func main() {
	http.HandleFunc(main_page, MainPage)
	http.HandleFunc(products_page, ProductsPage)
	http.HandleFunc(style, StyleCSS)
	http.HandleFunc(get_products_script, GetProductScript)
	http.HandleFunc(api_product, ProductService)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
