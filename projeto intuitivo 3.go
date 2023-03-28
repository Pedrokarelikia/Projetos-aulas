package main

import "fmt"

func main() {
	// Criando um slice de strings vazio
	nomes := []string{}

	// Adicionando os nomes ao slice
	nomes = append(nomes, "João", "Maria", "José")

	// Imprimindo o slice
	fmt.Println(nomes)
}
