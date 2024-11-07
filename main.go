package main

import (	
	"fmt"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func main() {
	contaDoJunior := contas.ContaCorrente{Titular: "LouJunior", Saldo: 200}
	contaDaLais := contas.ContaCorrente{Titular: "Lais", Saldo: 300}

	fmt.Println(contaDoJunior.Titular, "com Saldo de", contaDoJunior.Saldo)
	fmt.Println(contaDaLais.Titular, "com Saldo de", contaDaLais.Saldo)

	fmt.Println(contaDaLais.Transferir(100, &contaDoJunior))

	fmt.Println(contaDoJunior.Titular, "com Saldo de", contaDoJunior.Saldo)
	fmt.Println(contaDaLais.Titular, "com Saldo de", contaDaLais.Saldo)

}
