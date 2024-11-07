package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJunior := ContaCorrente{titular: "Junior", saldo: 100} //não precisa ter uma ordem, podemos remover alguns campos sem problemas
	contaDaLais := ContaCorrente{"Lais", 456, 7891010, 200}       //dessa maneira, entende que são argumentos, então é necessário que todos os campoes estejam preenchidos, caso contrário use a primeira forma, delacrando a variável
	fmt.Println(contaDoJunior, "\n", contaDaLais)

	//outra forma de criar os ponteiros, através do new, a estrutura acima é mais a cara de Go
	var contaDaMaria *ContaCorrente
	contaDaMaria = new(ContaCorrente)
	contaDaMaria.titular = "Maria"
	contaDaMaria.saldo = 500

	fmt.Println(contaDaMaria)

}
