package main

import "fmt"

func main() {
	var name = "William"
	var age int32 = 37
	const isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("The type of the value is %T", isCool)
}
