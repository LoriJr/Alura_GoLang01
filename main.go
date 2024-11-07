package main

import (
	"fmt"

	"github.com/LoriJr/Alura_GoLang01/clientes"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func main() {
	clienteJunior := clientes.Titular{Nome:"Junior", CPF:"123.123.123-12",Profissao: "Dev",}
	contaDoJunior := contas.ContaCorrente{Titular: clienteJunior,NumeroAgencia: 123,NumeroConta:123456, Saldo: 1000}
    fmt.Println(contaDoJunior)
}
