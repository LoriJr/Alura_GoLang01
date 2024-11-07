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

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64){
	if valorDaTransferencia <= c.saldo  && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia
		//contaDestino.saldo += valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso", c.saldo
	} else {
		return "Saldo insuficiente", c.saldo
	}

}

func main() {
contaDoJunior := ContaCorrente{titular:"LouJunior", saldo: 200}
contaDaLais := ContaCorrente{titular:"Lais", saldo: 300}

fmt.Println(contaDoJunior.titular, "com saldo de", contaDoJunior.saldo)
fmt.Println(contaDaLais.titular, "com saldo de", contaDaLais.saldo)

fmt.Println(contaDaLais.Transferir(100, &contaDoJunior))

fmt.Println(contaDoJunior.titular, "com saldo de", contaDoJunior.saldo)
fmt.Println(contaDaLais.titular, "com saldo de", contaDaLais.saldo)



}
