package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/todo/getAll", getAll)
	http.HandleFunc("/todo/getByID/", getById)
	http.HandleFunc("/todo/deleteByID/", deleteById)
	http.HandleFunc("/todo/addItem", addNode)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
