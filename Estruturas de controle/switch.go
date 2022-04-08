package main

import "fmt"

func notaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8,7:
		return "B"
	case 6,5:
		return "C"
	default:
		return "nota invalida"
	}
}

func main() {
	fmt.Println(notaConceito(7))
	fmt.Println(notaConceito(11))
	fmt.Println(notaConceito(2))
}