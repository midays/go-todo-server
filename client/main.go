package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var listFlag bool

func main() {
	var rootCMD = &cobra.Command{
		Use:   "todo-client",
		Short: "A simple CLI client to interact with the server",
		Run: func(cmd *cobra.Command, args []string) {
			if listFlag {

				req, err := http.NewRequest("GET", "http://127.0.0.1:8080/todo/getAll/", nil)
				if err != nil {
					log.Fatalf("Failed to create request: %v", err)
				}

				client := &http.Client{}
				response, err := client.Do(req)
				if err != nil {
					log.Fatalf("Failed to fetch todos: %v", err)
				}

				defer response.Body.Close()

				body, err := ioutil.ReadAll(response.Body)

				if err != nil {
					log.Fatalf("Failed to read response body: %v", err)
				}

				fmt.Println(string(body))

			}
		},
	}

	rootCMD.Flags().BoolVar(&listFlag, "list", false, "List all todo items")

	if err := rootCMD.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
