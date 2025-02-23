package main

import (
	"errors"
	"fmt"
)

func sumar(num1 int, num2 int) int {
	return num1 + num2
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir para 0")
	}

	return a / b, nil
}

// Ejemplo 1: Función para Invertir un Slice
// Una función que toma un slice de enteros y retorna otro slice con los elementos en orden inverso.

func invertirSlice(slice []int) []int {
	n := len(slice)
	invertido := make([]int, n)
	for i, v := range slice {
		invertido[n-i-1] = v
	}
	return invertido
}

func main() {
	// result := sumar(4, 5)
	// fmt.Println("La suma de 4 y 5 es:", result)
	//
	// result, err := dividir(10, 2)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("Resultado:", result)
	// }
	// Funciones anonimas

	// sumaFunc := func(a, b int) int {
	// 	return a + b
	// }

	// fmt.Println("Resultado:", sumaFunc(4, 3))

	// contador := func() func() int {
	// 	suma := 0
	// 	return func() int {
	// 		suma++
	// 		return suma
	// 	}
	// }()
	//
	// fmt.Println("Contador:", contador()) // Imprime: 1
	// fmt.Println("Contador:", contador()) // Imprime: 2
	// fmt.Println("Contador:", contador()) // Imprime: 3

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Numeros:", numeros)

	numerosInvertidos := invertirSlice(numeros)
	fmt.Println("Numeros Invertidos:", numerosInvertidos)
}
