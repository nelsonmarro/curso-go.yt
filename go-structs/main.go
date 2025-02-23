package main

import "fmt"

type Persona struct {
	Nombre    string
	Edad      int
	Altura    float32
	EsSoltero bool
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

func (p Persona) EsMayorDeEdad() bool {
	return p.Edad >= 18
}

func main() {
	nelson := Persona{
		Nombre:    "Nelson",
		Edad:      26,
		Altura:    1.70,
		EsSoltero: true,
	}

	// poner valores a los campos del struct
	nelson.Edad = 16
	nelson.Altura = 1.68

	nelson.Saludar()

	if nelson.EsMayorDeEdad() {
		fmt.Println("Es Mayor de Edad. Puede tomar Cerveza")
	} else {
		fmt.Println("Es Menor de Edad. Negado el alcohol!")
	}

	// Estamos leyendo los campos del struct
	// fmt.Printf("Hola %s Bienvenido!\n", nelson.Nombre)
	// fmt.Printf("Su edad es: %d\n", nelson.Edad)
	// fmt.Printf("Su altura es: %f\n", nelson.Altura)
	//
	// if nelson.EsSoltero {
	// 	fmt.Println("Es Soltero")
	// } else {
	// 	fmt.Println("Es Casado")
	// }

	// Declaración de un struct sin inicializar (los campos tendrán su valor cero):
	// var p3 Persona
}
