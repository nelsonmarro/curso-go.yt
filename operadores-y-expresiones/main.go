package main

import "fmt"

func main() {
	// Operadores Aritmeticos + - * / %
	var cantidad float64 = 3.0
	precio_producto := 20.5

	num1 := 20
	num2 := 2

	contador := 10

	total := cantidad * precio_producto
	suma := cantidad + precio_producto
	division := precio_producto / cantidad
	restoDivision := num1 % num2

	contador += 3
	fmt.Println("Contador:", contador)
	contador += 3
	fmt.Println("Contador:", contador)

	fmt.Println("El total de la venta es:", total)
	fmt.Println("El total de la suma es:", suma)
	fmt.Println("El total de la division es:", division)
	fmt.Println("El resto de la division es:", restoDivision)

	// Expresiones
	formula := ((3 + 5) * 10) / 2
	fmt.Println("El resultado de la formula es:", formula)
}
