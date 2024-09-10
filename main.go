package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	list    = List{}
	postsMu sync.Mutex
)

func getAll(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		http.Error(writer, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(list)

	postsMu.Lock()
	defer postsMu.Unlock()

}

func getById(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		http.Error(writer, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	id := request.URL.Path[len("/todo/getByID/"):]

	/*
		I couldn't do the compare on line 58, I had to change the foundNode to a pointer,
		and because of that I had to also change the code at line  54 I needed to add the '&' char. Why ?
	*/
	var foundNode *Node

	for _, node := range list.Nodes {

		if node.ID == id {
			foundNode = &node
		}
	}

	if foundNode != nil {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(foundNode)
	} else {
		http.Error(writer, "Item not found", http.StatusNotFound)
	}

	postsMu.Lock()
	defer postsMu.Unlock()

}

func addNode(writer http.ResponseWriter, response *http.Request) {

	if response.Method != "POST" {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
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

	list.Add(node)
	json.NewEncoder(writer).Encode("Item Added")

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

func main() {

	http.HandleFunc("/todo/getAll", getAll)
	http.HandleFunc("/todo/getByID/", getById)
	http.HandleFunc("/todo/deleteByID/", deleteById)
	http.HandleFunc("/todo/addItem", addNode)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
