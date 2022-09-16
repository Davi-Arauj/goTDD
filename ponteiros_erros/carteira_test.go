package ponteiroserros

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmarSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirada", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(10)}
		erro := carteira.Sacar(Bitcoin(5))
		confirmarSaldo(t, carteira, Bitcoin(5))
		confirmarErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Sacar(Bitcoin(100))

		confirmarSaldo(t, carteira, saldoInicial)
		confirmarErro(t, erro, ErroSaldoInsuficiente)

	})
}

func confirmarSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperdao '%s'", resultado, esperado)
	}
}

func confirmarErro(t *testing.T, resultado error, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Errorf("esperava um erro, mas nenhum ocorreu.")
	}

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

func confirmarErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}
