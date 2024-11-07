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
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func main() {
	contaJunior := ContaCorrente{}
	contaJunior.titular = "Junior"
	contaJunior.saldo = 1000
	
	fmt.Println(contaJunior.Sacar(200))
	fmt.Println(contaJunior.saldo)

}
