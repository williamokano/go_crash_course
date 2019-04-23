package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("At position %d the ID is %d\n", i, id)
	}

	// Not using id
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("The sum is ", sum)

	// Range with map
	emails := map[string]string{
		"Bob":      "bob@gmail.com",
		"Sharon":   "sharon@gmail.com",
		"Deleteme": "deleteme",
	}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: ", key)
	}

	for _, value := range emails {
		fmt.Println("Email is: ", value)
	}
}
