package main

import "fmt"

// pipeline
// Конвейер
func pipeline(in <-chan uint8, out chan<- float64) {
	go func() {
		for v := range in {
			num := float64(v) * float64(v) * float64(v)
			out <- num
		}
		close(out)
	}()
}

func main() {
	inChain := make(chan uint8)
	outChain := make(chan float64)
	numbers := []uint8{1, 2, 3, 4, 5}

	pipeline(inChain, outChain)

	go func() {
		for _, num := range numbers {
			inChain <- num
		}
		close(inChain)
	}()

	for result := range outChain {
		fmt.Println(result)
	}
}
