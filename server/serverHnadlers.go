package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	postsMu sync.Mutex
)

// todo: use gin + sqlite ( need GORM which gives the mapping between the DB objects, and the code )

func getAll(writer http.ResponseWriter, request *http.Request) {

	fmt.Print("\nGetting All items")
	postsMu.Lock()
	defer postsMu.Unlock()

	if request.Method != "GET" {
		http.Error(writer, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(list)

}

func getById(writer http.ResponseWriter, request *http.Request) {

	postsMu.Lock()
	defer postsMu.Unlock()

	if request.Method != "GET" {
		http.Error(writer, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	id := request.URL.Path[len("/todo/getByID/"):]

	node, found := list.getNodeByID(id)

	if found {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(node)
	} else {
		http.Error(writer, "Item not found", http.StatusNotFound)
	}

}

func addNode(writer http.ResponseWriter, response *http.Request) {

	postsMu.Lock()
	defer postsMu.Unlock()

	if response.Method != "POST" {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(writer, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	var node Node

	// Decode the JSON into the Node struct
	err = json.Unmarshal(body, &node)
	if err != nil {
		http.Error(writer, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	node.ID = generateID()

	result := db.Create(&node)
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(node)

	loadData()
}

func deleteById(writer http.ResponseWriter, request *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	id := request.URL.Path[len("/todo/deleteByID/"):]

	msg, deleted := list.Delete(id)

	if deleted {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("Item deleted")
	} else {
		http.Error(writer, msg, http.StatusNotFound)
	}
}
