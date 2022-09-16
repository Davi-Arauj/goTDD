package ponteiroserros

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")


func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
func (c *Carteira) Sacar(valor Bitcoin) error {
	if c.saldo < valor {
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
