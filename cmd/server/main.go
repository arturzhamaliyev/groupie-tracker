package main

import (
	"fmt"

	"groupie_tracker/internal"
)

func main() {
	fmt.Println("Starting server at http://localhost:8080")
	internal.Server()
}
