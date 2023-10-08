package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const MAXBUFSIZE = 4 * 1024 // 1Kb

type Header map[string]string

func errorHandler(w http.ResponseWriter, r *http.Request, url string, status int) int {
	if r.URL.String() == url {
		return 0
	}

	w.WriteHeader(status)
	w.Header().Add("Connection", "close")
	w.Write([]byte{})
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}

	return 1
}

func RespondPage(filename *string, h *Header, w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, MAXBUFSIZE)
	response_body := make([]byte, 0, MAXBUFSIZE)
	for cnt, err := file.Read(buf); cnt > 0; cnt, err = file.Read(buf) {
		if err != nil {
			panic(err)
		}

		response_body = append(response_body, buf[:cnt]...)
	}

	fmt.Printf("--------------------------------------------------------------\n%s %s\n", r.Method, r.URL)
	for key, val := range r.Header {
		fmt.Printf("%s: %s\n", key, val[0])
	}
	fmt.Printf("\n")

	for key, val := range *h {
		w.Header().Add(key, val)
	}
	w.Write(response_body)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if err := errorHandler(w, r, main_page, http.StatusNotFound); err != 0 {
		return
	}

	filename := "html/main.html"
	h := Header{
		"Content-Type": "text/html; charset=utf-8",
	}
	RespondPage(&filename, &h, w, r)
}

func ProductsPage(w http.ResponseWriter, r *http.Request) {
	if err := errorHandler(w, r, products_page, http.StatusNotFound); err != 0 {
		return
	}

	filename := "html/products.html"
	h := Header{
		"Content-Type": "text/html; charset=utf-8",
	}
	RespondPage(&filename, &h, w, r)
}

func StyleCSS(w http.ResponseWriter, r *http.Request) {
	if err := errorHandler(w, r, style, http.StatusNotFound); err != 0 {
		return
	}

	filename := "css/style.css"
	h := Header{
		"Content-Type": "text/css; charset=utf-8",
	}
	RespondPage(&filename, &h, w, r)
}

func GetProductScript(w http.ResponseWriter, r *http.Request) {
	if err := errorHandler(w, r, get_products_script, http.StatusNotFound); err != 0 {
		return
	}

	filename := "js/get_products.js"
	h := Header{
		"Content-Type": "application/javascript; charset=utf-8",
	}
	RespondPage(&filename, &h, w, r)
}

func ProductService(w http.ResponseWriter, r *http.Request) {
	if err := errorHandler(w, r, api_product, http.StatusNotFound); err != 0 {
		return
	}

	const table_name = "t_product"
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	type Product struct {
		Code        string
		Prod_name   string
		Weight      int
		Description string
	}

	query := "select code, prod_name, weight, description from " + table_name
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	http_res := make([]byte, 0, MAXBUFSIZE)
	http_res = append(http_res, []byte("{\"products\":[")...)

	for rows.Next() {
		row := Product{}
		if err := rows.Scan(&row.Code, &row.Prod_name, &row.Weight, &row.Description); err != nil {
			fmt.Println(err)
			continue
		}

		json_format, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}

		http_res = append(http_res, json_format...)
		http_res = append(http_res, ',')
	}
	http_res[len(http_res)-1] = ' '
	http_res = append(http_res, []byte("]}")...)
	fmt.Printf("%v\n", string(http_res))

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.Write(http_res)
}
