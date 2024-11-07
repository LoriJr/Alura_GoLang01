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
	Sacar(valor float64) string //como pagar um boleto é o mesmo que sacar um valor da conta, então nesse ponto o método sacar foi usado pq é a mesma idéia, e já está implementado, então não precisa fazer novamente essa parte do código.
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
