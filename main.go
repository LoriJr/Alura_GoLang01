package main

import (
	"fmt"
	//"github.com/LoriJr/Alura_GoLang01/clientes"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
contadoJunior := contas.ContaPoupanca {}
contadoJunior.Depositar(100)

contaDoPai := contas.ContaCorrente {}
contaDoPai.Depositar(500)
PagarBoleto(&contadoJunior, 60)
PagarBoleto(&contaDoPai, 400)

fmt.Println(contadoJunior.ObterSaldo())
fmt.Println(contaDoPai.ObterSaldo())

}
