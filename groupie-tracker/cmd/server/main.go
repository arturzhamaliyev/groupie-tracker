package main

import (
	"fmt"

	"groupie_tracker/internal"
)

func main() {
	fmt.Println("Starting server at http://localhost:1337")
	if err := internal.Server(); err != nil {
		fmt.Println(err)
		return
	}
}
