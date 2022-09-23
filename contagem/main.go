package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	ultimaPalavra  = "Vai!"
	inicioContagem = 3
)

// Sleeper te permite definir pausas
type Sleeper interface {
	Pausa()
}

// SleeperConfiguravel é uma implementação de Sleeper com uma pausa definida
type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

// Pausa
func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

// Pausa vai pausar a execução pela Duração definida
func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

// Contagem imprime uma contagem de 3 para a saída com um atraso determinado por um Sleeper
func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprintf(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
