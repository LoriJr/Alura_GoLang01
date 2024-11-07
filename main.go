package main

import (
	"fmt"

	"github.com/LoriJr/Alura_GoLang01/clientes"
	"github.com/LoriJr/Alura_GoLang01/contas"
)

func main() {
	contaDoJunior := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:"LouJr",
		CPF: "12345678912",
		Profissao: "analista de sistemas"},
		NumeroAgencia: 123,
		NumeroConta: 123456,
		Saldo:30}
	fmt.Println(contaDoJunior)
	//impressão fica assim {{LouJr 12345678912 analista de sistemas} 123 123456 30}
	//são duas chaves, a primeira se refere aos dados da conta corrente, quanto a segunda se refere aos dados do titular, com nome cpf e profissão

}
