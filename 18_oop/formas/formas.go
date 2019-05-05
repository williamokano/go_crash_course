package formas

import "math"

// Retangulo whatever
type Retangulo struct {
	Base   float64
	Altura float64
}

// Area asdas
func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}

// Circulo asdasd
type Circulo struct {
	Raio float64
}

// Area asas
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
