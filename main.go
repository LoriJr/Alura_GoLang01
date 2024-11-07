package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	
	var contaDaMaria *ContaCorrente
	contaDaMaria = new(ContaCorrente)
	contaDaMaria.titular = "Maria"
	contaDaMaria.saldo = 500

	var contaDaMaria2 *ContaCorrente
	contaDaMaria2 = new(ContaCorrente)
	contaDaMaria2.titular = "Maria"
	contaDaMaria2.saldo = 500

	fmt.Println(contaDaMaria == contaDaMaria2) // nessa comparação mesmo que os dados sejam iguais, a comparação é do endereço na memória, que no caso é falso
	fmt.Println(*contaDaMaria == *contaDaMaria2) // estamos comparando os dados das variáveis, que no caso são iguais e portanto, verdadeiro
	fmt.Println(&contaDaMaria)  //endereço na memória 0xc00005e050 
	fmt.Println(&contaDaMaria2) //endereço na memória 0xc00005e058
	
	
	
}
