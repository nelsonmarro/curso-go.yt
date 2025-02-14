package main

import "fmt"

func main() {
	// Arrays
	// var nombres [3]string = [3]string{"Nelson", "Pablo", "Jose"}
	// animales := [3]string{"Perro", "Gato", "Jirafa"}
	// fmt.Println(nombres)
	// fmt.Println(animales)

	// Acceder a un elemento especifico del array
	// fmt.Println(animales[1])

	// Cambiar un valor de un elemento del array
	// animales[0] = "Leon"
	// fmt.Println(animales[0])

	// Slices
	numeros := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numeros)

	// Agregar elementos al final del slice
	numeros = append(numeros, 8)
	fmt.Println(numeros)

	promedios := make([]float32, 0, 10)
	promedios = append(promedios, 8.5)
	promedios = append(promedios, 9.1)
	promedios = append(promedios, 2.5)
	fmt.Println(len(promedios))

	// Copiar el valor de un slice a otro
	promediosCopy := make([]float32, len(promedios))
	copy(promediosCopy, promedios) // Copia los elementos de promedios a promediosCopy
	fmt.Println("Valor de promedios:", promedios)
	fmt.Println("Valor de promedios Copy:", promedios)
}
