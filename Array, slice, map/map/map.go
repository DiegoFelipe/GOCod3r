package main

import "fmt"

// Map é como um obj em javascript

func main() {
	// var aprovados map[int]string
	// maps devem ser inicialidados

	aprovados := make(map[int]string)

	aprovados[732676734] = "Diego"
	aprovados[8732472347823478] = "Felipe"
	aprovados[010234234234234] = "Souza"

	fmt.Println(aprovados)

	// key, value
	/**
		Diego (CPF: 732676734)
		Felipe (CPF: 8732472347823478)
		Souza (CPF: 570734753948)
	*/
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome,cpf)
	}

	// acessar um especifico
	fmt.Println(aprovados[8732472347823478]) // Felipe

	// para deletar
	delete(aprovados,8732472347823478)
	fmt.Println(aprovados)
}