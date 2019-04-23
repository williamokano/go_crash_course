package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello " + p.firstName + " " + p.lastName + " of age " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouse Person) {
	if p.gender == "M" {
		return
	}
	p.lastName = spouse.lastName
}

func main() {
	person1 := Person{
		firstName: "Samanta",
		lastName:  "Quux",
		city:      "New York",
		gender:    "F",
		age:       25,
	}

	person2 := Person{
		firstName: "Bob",
		lastName:  "Okano",
		city:      "Uberlândia",
		gender:    "M",
		age:       30,
	}

	person1.getMarried(person2)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
