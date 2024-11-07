package main

import (
	"fmt"

	"github.com/LoriJr/Alura_GoLang01/clientes"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func main() {
	contaDoJunior := contas.ContaCorrente{Titular: clientes.Titular{Nome:"LouJr", CPF: "32055566644488", Profissao: "analista de sistemas"}, Saldo:30}
	contaDaLais := contas.ContaCorrente{Titular: clientes.Titular{Nome:"Lais"}, Saldo: 300}

	fmt.Println(contaDoJunior.Titular, "com Saldo de", contaDoJunior.Saldo)
	fmt.Println(contaDaLais.Titular, "com Saldo de", contaDaLais.Saldo)

	fmt.Println(contaDaLais.Transferir(100, &contaDoJunior))

	fmt.Println(contaDoJunior.Titular, "com Saldo de", contaDoJunior.Saldo)
	fmt.Println(contaDaLais.Titular, "com Saldo de", contaDaLais.Saldo)

}
