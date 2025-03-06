package sumaconcurrente

func Sumar(numeros []int, canalResults chan int) {
	suma := 0
	for _, num := range numeros {
		suma += num
	}
	canalResults <- suma
}
