package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	fmt.Printf("Hola, %s desde una goroutine!\n", nombre)
}

func main() {
	go saludar("Gophers!")
	fmt.Println("Hola desde la funcion principal!")
	time.Sleep(1 * time.Second)
}
