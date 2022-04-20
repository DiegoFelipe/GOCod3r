package main

import "fmt"

func main() {
	// um map com um map dentro dele
	funsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Sila": 1212,
			"Guga":          19822,
		},
		"D": {
			"Diego Felipe": 48000,
		},
	}

	for letra, funcs := range funsPorLetra {
		/**
		        G map[Gabriela Sila:1212 Guga:19822]
				D map[Diego Felipe:48000]
		*/
		fmt.Println(letra)

		for nome, salario := range funcs {
			fmt.Printf("%v ganha %.2f \n", nome, salario)
		}
	}
}
