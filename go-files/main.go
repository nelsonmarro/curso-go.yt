package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rutaArchivo := "notas.txt"
	archivoNotas, err := os.Open(rutaArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer archivoNotas.Close()

	fmt.Printf("Contenido de '%s':\n", rutaArchivo)
	escaner := bufio.NewScanner(archivoNotas)

	// Leer línea a línea
	for escaner.Scan() {
		linea := escaner.Text()
		fmt.Println(linea)
	}
}
