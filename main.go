package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJunior := ContaCorrente{titular: "Junior", saldo: 100} 
	contaDoJunior2 := ContaCorrente{titular: "Junior", saldo: 100} 
	fmt.Println(contaDoJunior) //aqui o código entende que a variável pertence à estrura <contaCorrente>
	fmt.Println(&contaDoJunior) //aqui o código entende que a variável contém as mesmas informações da estrutura, mas não QUE SEJA da mesma estrutura
	fmt.Println(&contaDoJunior == &contaDoJunior2) // com o &, está comparando se os endereços são iguais, e no caso é falso
	fmt.Println(contaDoJunior == contaDoJunior2) // está comparando o conteúdo das variáveis, e no caso é verdadeiro

	//abaixo é feito o mesmo comparativo, mas nos campos que são como argumentos, e no caso a comparação é igual a de cima
	contaDaLais := ContaCorrente{"Lais", 456, 7891010, 200}
	contaDaLais2 := ContaCorrente{"Lais", 456, 7891010, 200}       
	fmt.Println(contaDaLais == contaDaLais2)	
}
