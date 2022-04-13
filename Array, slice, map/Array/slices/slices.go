package main

import "fmt"

func main() {
	a1 := [3]int{1,2,3} // array
	slice1 := []int{1,2,3} // slice

	fmt.Println(a1,slice1)

	a2 := [5]int{1,2,3,4,5} // 

	// slice não é um array, é um pedaço de um array
	s2 := a2[1:3] // até o index 3, mas não inclui esse index
	fmt.Println(a2,s2) // [2,3]

	s3 := a2[:2] // [1,2]
	fmt.Println(s3)

	// slice = tamanho, pointer pra um array, para uma parte do array

	// pode ter slice de um slice
	sliceDoSlice := s3[:1]
	fmt.Println(sliceDoSlice) // [1]
}