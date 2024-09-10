package main

import "github.com/google/uuid"

func generateID() string {
	uuid := uuid.New()
	return uuid.String()
}
