package main

import (
	"fmt"
	"os"

	"gopkg.in/resty.v1"
)

// Todo whatever
type Todo struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

func main() {
	resp, error := resty.R().
		SetResult(&[]Todo{}).
		Get("https://jsonplaceholder.typicode.com/todos")

	if error != nil {
		fmt.Println("Faio", error)
		os.Exit(1)
	}

	todoList := *resp.Result().(*[]Todo)

	for _, todoItem := range todoList {
		fmt.Printf("ID: %d\n", todoItem.ID)
	}
}
