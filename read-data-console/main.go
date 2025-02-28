package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	// var nombre string
	// var edad int
	//
	// fmt.Print("Introduce to Nombre y tu Edad (separados por espacios): ")
	// fmt.Scanf("%s , %d", &nombre, &edad)
	//
	// fmt.Println("Nombre:", nombre)
	// fmt.Println("Edad:", edad)

	// var nombreCompleto string
	//
	// fmt.Print("Introduce tu nombre completo: ")
	// fmt.Scanln(&nombreCompleto)
	//
	// fmt.Println("Nombre Completo:", nombreCompleto)
	//

	// reader := bufio.NewReader(os.Stdin)
	//
	// fmt.Print("Introduce tu mensaje: ")
	//
	// mensaje, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error al leer:", err)
	// 	return
	// }
	//
	// mensaje = strings.TrimSuffix(mensaje, "\n")
	// fmt.Println("Mensaje recibido:", mensaje)
	//

	reader := bufio.NewReader(os.Stdin)
	persona := Persona{}

	fmt.Print("Introduce tu nombre: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error al leer nombre: %w", err)
	}
	persona.Nombre = strings.TrimSuffix(nombre, "\n")
	fmt.Println(persona.Nombre)
}
