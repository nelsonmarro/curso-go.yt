package main

import (
	"fmt"

	sumaconcurrente "github.com/nelsonmarro/go-concurrency/suma_concurrente"
)

// func saludar(nombre string) {
// 	fmt.Printf("Hola, %s desde una goroutine!\n", nombre)
// }

func main() {
	// go saludar("Gophers!")
	// fmt.Println("Hola desde la funcion principal!")
	// time.Sleep(1 * time.Second)
	numeros1 := []int{1, 2, 3, 4, 5}
	numeros2 := []int{6, 7, 8, 9, 10}

	// declaración y creación de un canal para la suma concurrente
	canalSuma := make(chan int)

	// creación de dos goroutines para sumar los números del slice concurrentemente
	go sumaconcurrente.Sumar(numeros1, canalSuma)
	go sumaconcurrente.Sumar(numeros2, canalSuma)

	// esperar a que las goroutines terminen y obtener la suma total
	suma1 := <-canalSuma
	suma2 := <-canalSuma

	fmt.Printf("Suma 1: %d, Suma 2: %d, Suma Total: %d\n", suma1, suma2, suma1+suma2)
}
