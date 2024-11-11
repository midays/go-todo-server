package main

type Node struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Priority int    `json:"priority,omitempty"`
}
