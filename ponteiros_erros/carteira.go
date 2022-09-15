package ponteiroserros

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface{
	String() string
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
