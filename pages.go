package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const MAXBUFSIZE = 4 * 1024 // 1Kb

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.Write([]byte("Body data"))
}

func ProductService(w http.ResponseWriter, r *http.Request) {
	const table_name = "t_product"
	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	type Product struct {
		Id          int
		Code        string
		Prod_name   string
		Weight      int
		Description string
	}

	query := "select * from " + table_name
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	http_res := make([]byte, 0, MAXBUFSIZE)

	for rows.Next() {
		row := Product{}
		if err := rows.Scan(&row.Id, &row.Code, &row.Prod_name, &row.Weight, &row.Description); err != nil {
			fmt.Println(err)
			continue
		}

		json_format, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}

		http_res = append(http_res, json_format...)
	}
	fmt.Printf("%v\n", string(http_res))

	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.Write(http_res)
}
