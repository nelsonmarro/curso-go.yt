package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

// Funcion para crear una nueva instancia de la estructura Persona
func NewPersona(nombre string, edad int) *Persona {
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

func (p *Persona) CumplirAniosPunteros() {
	if p != nil {
		p.Edad++
		fmt.Printf("¡Feliz cumpleaños, %s! Ahora tienes %d años (modificado con puntero).\n", p.Nombre, p.Edad)
	} else {
		fmt.Println("No se proporcionó una persona válida para celebrar el cumpleaños.")
	}
}

func CumplirAniosSinPunteros(persona Persona) {
	persona.Edad++
	fmt.Printf("¡Feliz cumpleaños, %s! Ahora tienes %d años (modificado sin puntero).\n", persona.Nombre, persona.Edad)
}

func main() {
	ana := &Persona{Nombre: "Ana", Edad: 25}
	pablo := NewPersona("Pablo", 30)
}
