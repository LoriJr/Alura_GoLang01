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
	if podeSacar { //se podeSacar for verdadeiro
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (d *ContaCorrente) Deposito(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar { // se podeDepositar for verdadeiro
		d.saldo += valorDoDeposito
		return "Deposito realizado com sucesso"
	} else {
		return "Valor do deposito menor ou igual a zero"
}
}

func main() {
	contaJunior := ContaCorrente{}
	contaJunior.titular = "Junior"
	contaJunior.saldo = -100
	
	fmt.Println(contaJunior.Deposito(200))
	fmt.Println(contaJunior.saldo)

}
