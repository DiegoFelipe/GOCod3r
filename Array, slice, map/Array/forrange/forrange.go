package main

import "fmt"

func main() {
	numeros := [...]int{1,2,3,4,5,6,7,8}

	// for _ => para ignorar o index
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}
}