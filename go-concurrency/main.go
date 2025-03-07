package main

import (
	"fmt"

	"github.com/nelsonmarro/go-concurrency/opconcurrentes"
)

// func saludar(nombre string) {
// 	fmt.Printf("Hola, %s desde una goroutine!\n", nombre)
// }

func main() {
	// invocar una funcion en una goroutine (de manera concurrente)
	// go saludar("CodeCrafter")
	// fmt.Println("Ejecutando en la goroutine principal...")
	// time.Sleep(1 * time.Second)
	numeros1 := []int{1, 2, 3, 4, 5}
	numeros2 := []int{6, 7, 8, 9, 10}

	// canal sin buffer
	canalSumas := make(chan int)

	// llamar a la funcion Sumar de manera concurrente
	go opconcurrentes.Sumar(numeros1, canalSumas)
	go opconcurrentes.Sumar(numeros2, canalSumas)

	for suma := range canalSumas {
		fmt.Printf("Suma: %d\n", suma)
	}

	// fmt.Printf("Suma 1: %d, Suma 2: %d, Suma Total: %d\n", suma1, suma2, suma1+suma2)
}
