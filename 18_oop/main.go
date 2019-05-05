package main

import (
	"fmt"

	"github.com/williamokano/go_crash_course/18_oop/formas"
)

func printArea(forma interface{ Area() float64 }) {
	area := forma.Area()
	fmt.Printf("A area é: %f\n", area)
}

func main() {
	circulo := formas.Circulo{Raio: 10}
	retangulo := formas.Retangulo{Base: 5, Altura: 10}

	printArea(circulo)
	printArea(retangulo)
}
