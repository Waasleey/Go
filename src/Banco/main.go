package main

import (
	clientes "Banco/clientes"
	contas "Banco/contas"
	"fmt"
)

func main() {

	contadoWasley := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:         "Wasley",
			Sobrenome:    "Fernando",
			NomeExibicao: "Waasleey",
			TitularCPF:   "111.111.111.10",
			Profissao:    "Analista",
		},
		NumeroAgencia: 111,
		NumeroConta:   225110,
	}

	fmt.Println(contadoWasley.ExibeSaldo())

	contadoWasley.Depositar(200)
	fmt.Println(contadoWasley.ExibeSaldo())

}
