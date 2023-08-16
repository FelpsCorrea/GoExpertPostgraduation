package main

type Conta struct {
	saldo int
}

// Retorna o endereço da nova conta
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// Altera o valor do saldo
func (c *Conta) deposito(valor int) {
	c.saldo += valor
}

// Não altera o valor do saldo
func (c Conta) simularDeposito(valor int) {
	c.saldo += valor
}

func main() {

}
