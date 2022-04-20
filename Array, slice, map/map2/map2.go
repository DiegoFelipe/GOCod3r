package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":    121212,
		"Diego Felipe": 50000,
	}

	funcsESalarios["Guga"] = 97000
	delete(funcsESalarios, "Funcionario inecxistente") // Não dá erro se tentar deletar um valor que não existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
