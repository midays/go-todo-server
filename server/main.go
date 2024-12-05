package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	list = List{}
)

func initializeDatabase() {
	var err error

	db, err = gorm.Open(sqlite.Open("../assets/todo.db"))
	if err != nil {
		log.Fatal("Failed to open the database file")
	}
	db.AutoMigrate(&Node{})
	loadData()
}

func loadData() {
	var nodes []Node
	result := db.Find(&nodes)
	if result.Error != nil {
		log.Println("Error loading data from database:", result.Error)
	}
	list.Nodes = nodes
	log.Printf("Loaded %d todo items from the database.\n", len(list.Nodes))
}

func main() {

	initializeDatabase()

	http.HandleFunc("/todo/getAll", getAll)
	http.HandleFunc("/todo/getByID/", getById)
	http.HandleFunc("/todo/deleteByID/", deleteById)
	http.HandleFunc("/todo/addItem", addNode)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
