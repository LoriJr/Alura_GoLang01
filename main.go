package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo //podeSacar é um booleano, que aceita verdadeiro ou falso
	if podeSacar {                                           //se podeSacar for verdadeiro
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar { // se podeDepositar for verdadeiro
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor ou igual a zero", c.saldo
	}
}

func main() {
	contaJunior := ContaCorrente{}
	contaJunior.titular = "Junior"
	contaJunior.saldo = 100

	// primeira forma
	fmt.Println(contaJunior.saldo)
	fmt.Println(contaJunior.Depositar(200))

	//segunda forma
	//trabalhando com os returns abaixo
	status, valor := contaJunior.Depositar(200) // é como se fosse um unpacking do python
	fmt.Println(status, valor)

}
