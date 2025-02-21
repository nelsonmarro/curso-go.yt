package main

import "fmt"

func main() {
	// for clasico
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Numero: %d\n", i)
	// }

	// for como while
	// palabras := []string{"hola", "mundo", "como", "estas"}
	// for palabras[0] == "hola" {
	// 	fmt.Printf("Primera Palabra: %s \n", palabras[0])
	// 	palabras[0] = "adios"
	// }
	// fmt.Printf("Primera Palabra: %s \n", palabras[0])

	// f := 1
	// num := 6
	//
	// for i := num; i >= 1; i-- {
	// 	f *= i
	// }
	//
	// fmt.Printf("Factorial de %d es: %d\n", num, f)

	// for _, valor := range palabras {
	// 	fmt.Printf("Palabra: %s \n", valor)
	// }

	edades := map[string]int{
		"Ana":    30,
		"Luis":   25,
		"Carlos": 40,
	}

	// Usando for range para iterar sobre el map
	for key, valor := range edades {
		fmt.Printf("Nombre: %s, Edad: %d\n", key, valor)
	}
}
