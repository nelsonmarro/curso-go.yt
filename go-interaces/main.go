package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func main() {
	c := Circulo{Radio: 5}
	r := Rectangulo{Base: 3, Altura: 4}

	// fmt.Printf("Area del circulo: %.2f\n", c.Area())
	// fmt.Printf("Área del rectángulo: %.2f\n", r.Area())

	var f1 Figura = r
	var f2 Figura = c

	fmt.Printf("Área de f2 (círculo): %.2f\n", f2.Area())
	fmt.Printf("Área de f1 (rectángulo): %.2f\n", f1.Area())
}
