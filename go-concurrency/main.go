package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	fmt.Printf("Hola, %s desde una goroutine!\n", nombre)
}

func main() {
	// invocar una funcion en una goroutine (de manera concurrente)
	go saludar("CodeCrafter")
	fmt.Println("Ejecutando en la goroutine principal...")
	time.Sleep(1 * time.Second)
}
