package main

import "fmt"

func main() {
	// Inicializacion literal de los maps
	elementosTablaPeriodica := map[string]string{
		"H":  "Hidrogeno",
		"Li": "Litio",
		"K":  "Potasio",
		"Ca": "Oxigeno",
	}
	fmt.Println(elementosTablaPeriodica)

	// edades_nombres := make(map[string]int)
	// edades_nombres["Nelson"] = 26
	// edades_nombres["Rafael"] = 13
	//
	// fmt.Println(edades_nombres)

	litio, ok := elementosTablaPeriodica["Li"]
	if ok {
		fmt.Println(litio)
	} else {
		fmt.Println("No se encuentra la clave mapa")
	}

	// actualizar un valor del mapa
	elementosTablaPeriodica["Ca"] = "Calcio"
	fmt.Println(elementosTablaPeriodica)

	// eliminar un elemento
	delete(elementosTablaPeriodica, "Ca")
	fmt.Println(elementosTablaPeriodica)
}
