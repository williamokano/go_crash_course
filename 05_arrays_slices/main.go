package main

import "fmt"

func main() {
	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)
	// fmt.Println(fruitArray[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])
}
