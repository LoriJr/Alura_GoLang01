package main

import (
	"fmt"
	//"github.com/LoriJr/Alura_GoLang01/clientes"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func main() {
contadoJunior := contas.ContaPoupanca {}
contadoJunior.Depositar(100)
fmt.Print(contadoJunior.ObterSaldo())

}
