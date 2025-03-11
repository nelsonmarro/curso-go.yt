package main

import (
	"bufio"
	"fmt"
	"io"
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

	// abrir el archivo creado para leerlo
	miarchivoLectura, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer miarchivoLectura.Close()

	fmt.Printf("\nContenido de '%s' (después de escritura inicial):\n", nombreArchivo)
	escanerInicial := bufio.NewScanner(miarchivoLectura) // Escáner inicial
	for escanerInicial.Scan() {                          // Leer líneas
		linea := escanerInicial.Text() // Obtener línea
		fmt.Printf("%s\n", linea)
	}
	if err := escanerInicial.Err(); err != nil { // Comprobar errores
		fmt.Println("Error al leer después de escritura inicial:", escanerInicial.Err())
	}

	// Append para agregar más texto al archivo
	archivoAppend, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error al abrir el archivo para append:", err)
	}
	defer archivoAppend.Close()

	textoAppend := "\n--- Texto añadido (APPEND) ---\nLínea añadida.\n¡Append OK!\n"
	fmt.Printf("\nAñadiendo texto (APPEND) a '%s':\n%s", nombreArchivo, textoAppend)

	_, err = io.WriteString(archivoAppend, textoAppend)
	if err != nil {
		fmt.Println("Error al escribir el texto (APPEND):", err)
	}

	// Eliminar un archivo
	err = os.Remove("notas.txt")
	if err != nil {
		fmt.Println("Error al eliminar el archivo:", err)
	}
	fmt.Printf("\nArchivo '%s' eliminado.\n", "notas.txt")
}
