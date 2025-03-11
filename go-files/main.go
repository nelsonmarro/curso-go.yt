package main

import (
	"fmt"
	"os"
)

// Leer un archivo existente y mostrar su contenido
// func main() {
// 	rutaArchivo := "notas.txt"
// 	archivoNotas, err := os.Open(rutaArchivo)
// 	if err != nil {
// 		fmt.Println("Error al abrir el archivo:", err)
// 	}
// 	defer archivoNotas.Close()
//
// 	fmt.Printf("Contenido de '%s':\n", rutaArchivo)
// 	escaner := bufio.NewScanner(archivoNotas)
//
// 	// Leer línea a línea
// 	for escaner.Scan() {
// 		linea := escaner.Text()
// 		fmt.Printf("%s\n", linea)
// 	}
//
// 	if err := escaner.Err(); err != nil {
// 		fmt.Println("Error al leer el archivo:", err)
// 	}
// }

func main() {
	nombreArchivo := "mi-archivo.txt"

	miarchivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	textoInicial := "Texto inicial simplificado.\nLínea uno.\nLínea dos.\n"
	fmt.Printf("Escribiendo texto inicial en '%s':\n%s", nombreArchivo, textoInicial)

	_, err = miarchivo.WriteString(textoInicial)
	if err != nil {
		fmt.Println("Error al escribir el texto:", err)
	}
}
