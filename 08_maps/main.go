package main

import "fmt"

func main() {
	emails := map[string]string{
		"Bob":      "bob@gmail.com",
		"Sharon":   "sharon@gmail.com",
		"Deleteme": "deleteme",
	}

	delete(emails, "deleteme")

	fmt.Println(emails)
}
