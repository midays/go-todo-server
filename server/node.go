package main

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Priority  int    `json:"priority,omitempty"`
	Completed bool   `json:"completed"`
}
