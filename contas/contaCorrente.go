package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo //podeSacar é um booleano, que aceita verdadeiro ou falso
	if podeSacar {                                           //se podeSacar for verdadeiro
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar { // se podeDepositar for verdadeiro
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do deposito menor ou igual a zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64){
	if valorDaTransferencia <= c.Saldo  && valorDaTransferencia > 0{
		c.Saldo -= valorDaTransferencia
		//contaDestino.Saldo += valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}

}