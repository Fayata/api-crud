package main

import (
	"api-crud/route"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		route.Router(rw, r)
	})

	fmt.Println("Menjalankan server di http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
