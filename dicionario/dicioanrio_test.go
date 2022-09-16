package dicionario

import (
	"testing"
)

func TestDeleta(t *testing.T) {
	t.Run("palavra deletada", func(t *testing.T) {
		palavra := "teste"
		dicionario := Dicionario{palavra: "definição teste"}

		dicionario.Deleta(palavra)

		_, err := dicionario.Busca(palavra)
		if err != ErrNaoEncontrado {
			t.Errorf("espera-se que '%s' seja deletado", palavra)
		}
	})
}

func TestAtualizar(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		novaDefinicao := "nova definição"

		err := dicionario.Atualizar(palavra, novaDefinicao)
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualizar(palavra, definicao)

		comparaErro(t, err, ErrPalavraInexistente)
	})

}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "é só um teste"

		err := dicionario.Adiciona(palavra, definicao)
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)

	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra2 := "teste"
		definicao2 := "é só um teste"

		err := dicionario.Adiciona(palavra2, definicao2)
		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra2, definicao2)
	})

}

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "é só um teste"}

	t.Run("palvara conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "é só um teste"

		comparaStrings(t, resultado, esperado)
	})
	t.Run("palvara desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")
		comparaErro(t, resultado, ErrNaoEncontrado)
	})

}

// comparaErro - compara dois erros
func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

// comparaStrings - compara duas strings
func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado: '%s', esperado: '%s', dado: '%s'", resultado, esperado, "teste")
	}
}

// comparaDefinicao - compara a definicao de uma palavra no dicionario
func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()
	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("não foi possivel encontrar a palvra adicionada: ", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
}
