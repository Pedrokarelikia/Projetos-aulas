package main

import "fmt"

func main() {
	// criando o slice de inteiros
	numeros := []int{1, 2, 3, 4, 5}

	// removendo o terceiro elemento
	numeros = append(numeros[:2], numeros[3:]...)

	// imprimindo o slice resultante
	fmt.Println(numeros)
}