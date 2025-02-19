package main

import "fmt"

func main() {
	// declarar una variable
	// edad := 8
	//
	// if edad >= 18 {
	// 	fmt.Println("Es mayor de edad")
	// } else {
	// 	fmt.Println("Es menor de edad")
	// }

	// if con declaración de variable incluida
	// if edad := 21; edad >= 18 {
	// 	fmt.Println("Es mayor de edad")
	// } else {
	// 	fmt.Println("Es menor de edad")
	// }

	// else if
	// calificacion := 100
	//
	// if calificacion >= 90 {
	// 	fmt.Println("Sobresaliente")
	// } else if calificacion >= 80 {
	// 	fmt.Println("Notable")
	// } else if calificacion >= 70 {
	// 	fmt.Println("Bien")
	// } else {
	// 	fmt.Println("Su calificacio es menor  de 70. Necesita mejorar")
	// }

	// switch
	dia := 199

	switch dia {
	case 1, 2, 3, 4, 5:
		fmt.Println("Es un día laborable")
	case 6, 7:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Día desconocido")
	}
}
