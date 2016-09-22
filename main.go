package main

import (
	"fmt"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!</h1>\n")
    for _, e := range os.Environ() {
        fmt.Fprintf(w, e)
    }
}

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
