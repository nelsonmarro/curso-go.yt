package opconcurrentes

func Sumar(numeros []int, canalResult chan int) {
	suma := 0

	for _, num := range numeros {
		suma += num
	}
	canalResult <- suma
}

func Restar(numeros []int, canalResult chan int) {
	resta := numeros[0]
	for _, num := range numeros {
		resta -= num
	}

	canalResult <- resta
}
