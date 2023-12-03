package contas

import clientes "Banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float32
}

func (c *ContaCorrente) Sacar(valorASacar float64) (string, float64) {

	if c.validaSaque(valorASacar) {
		c.saldo -= float32(valorASacar)
		return "Saque realizado com sucesso. saldo atual: ", float64(c.saldo)
	} else {
		return "saldo insuficiente para saque. saldo atual: ", float64(c.saldo)
	}
}

func (c *ContaCorrente) validaSaque(valorARemover float64) bool {

	if valorARemover > 0 && valorARemover <= float64(c.saldo) {
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Depositar(valorADepositar float64) (string, float64) {

	if valorADepositar > 0 {
		c.saldo += float32(valorADepositar)
		return "Deposito realizado com sucesso. saldo atual: ", float64(c.saldo)
	} else {
		return "Deposito incorreto. saldo atual: ", float64(c.saldo)
	}
}

func (c *ContaCorrente) Tranferir(contaDestino *ContaCorrente, valorATransferir float64) bool {

	if c.validaSaque(valorATransferir) {
		c.saldo -= float32(valorATransferir)
		contaDestino.Depositar(valorATransferir)
		return true
	} else {
		return false
	}
}
func (c *ContaCorrente) ExibeSaldo() float32 {
	return c.saldo
}
