package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)
	
	s = make([]int, 10, 20) // 10 elementos o array interno vai ter 20 posições
	fmt.Println(s, cap(s)) // cap(s) = the internal array capacity

	s = append(s,1,2,3,4,5,6,7,8,9,0)
	fmt.Println(s, len(s), cap(s)) // ate esse ponto a capacidade é 20

	// mas se adicionar mais um elemento no slice, ele aumenta de tamanho
	s = append(s,1) // a partir de agora o tamanho do slice é 21 e a capacidade 40
	fmt.Println(s, len(s), cap(s)) // dobrou a capacidade aqui, agora é 40 e o len 21
}