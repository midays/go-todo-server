package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(request.URL.Path[len("/todo/getByID/"):])
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

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

func AddNode(w http.ResponseWriter, r *http.Request, id int) {

}

func deleteById(writer http.ResponseWriter, request *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	id, err := strconv.Atoi(request.URL.Path[len("/todo/deleteByID/"):])
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	msg, deleted := list.Delete(id)

	if deleted {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("Item deleted")
	} else {
		http.Error(writer, msg, http.StatusNotFound)
	}
}

func main() {

	list.Add(Node{ID: 1, Name: "first node"})
	list.Add(Node{ID: 2, Name: "second node"})
	list.Add(Node{ID: 3, Name: "third node"})

	http.HandleFunc("/todo/getAll", getAll)
	http.HandleFunc("/todo/getByID/", getById)
	http.HandleFunc("/todo/deleteByID/", deleteById)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
