package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var input int

	fmt.Print("Insira o número máximo que deseja ser sorteado: ")
	fmt.Scan(&input)

	var numeros [10]int

	for posicao, valor := range numeros {

		valor = rand.Intn(input)
		fmt.Println("Na posição", posicao, "foi sorteado o número:", valor)
	}

}
