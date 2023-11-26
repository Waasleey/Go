package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float32
}

func main() {

	contadoWasley := ContaCorrente{

		titular:       "Wasley",
		numeroAgencia: 153,
		numeroConta:   225123,
		saldo:         254.90,
	}

	contadoFernando := ContaCorrente{
		titular:       "Fernando",
		numeroAgencia: 153,
		numeroConta:   2252371,
		saldo:         251.87,
	}
	fmt.Println(contadoWasley)
	fmt.Println(contadoFernando)
}
