package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "log"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-type", "text/html; charset=utf-8")
    w.Write([]byte("Body data"))
}


func ProductService(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-type", "application/json; charset=utf-8")
    
    type Resp struct {
        Id int
        Name string
    }
    
    resp := Resp {
        Id: 999,
        Name: "hello",
    }
    fmt.Printf("%v\n", resp)
    
    httpresp, err := json.Marshal(resp)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%v\n", httpresp)
    fmt.Printf("%v\n", string(httpresp))
    w.Write(httpresp)
}
