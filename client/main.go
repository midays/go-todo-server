package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	getByIDArg string
	listFlag   bool
	baseURL    = "http://127.0.0.1:8080/"
)

func main() {
	var rootCMD = &cobra.Command{
		Use:   "todo-client",
		Short: "A simple CLI client to interact with the server",
		// todo: switch to RunE, used to handle errors in a correct friendly way
		Run: func(cmd *cobra.Command, args []string) {

			if listFlag {
				url := baseURL + "todo/getAll"
				sendGetRequest(url)
			}

			if getByIDArg != "" {
				url := baseURL + "todo/getByID/" + getByIDArg
				sendGetRequest(url)
			}
			if !listFlag && getByIDArg == "" {
				fmt.Print("A simple CLI client to interact with the server\n")
				fmt.Print("Usage:\n")
				fmt.Print("\t--list - to list all todo items\n")
				fmt.Print("\t--getById <id> - to get a specific item\n")
			}
		},
	}

	rootCMD.Flags().BoolVar(&listFlag, "list", false, "List all todo items")
	rootCMD.Flags().StringVar(&getByIDArg, "getById", "", "Get a todo item by ID")

	if err := rootCMD.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func sendGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
}
