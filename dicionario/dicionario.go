package dicionario

type Dicionario map[string]string

type ErrDicionario string

// Error - recebe um erro do tipo ErrDicionario e retorna uma mensagem
func (e ErrDicionario) Error() string {
	return string(e)
}

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente   = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("não foi possível atualizar a palavra pois ela não existe")
)

// Deleta - apaga uma palavra do dicionario
func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}

// Atualizar - realiza uma atualização na definição de uma palavra no dicionario
func (d Dicionario) Atualizar(palavra, definicao string) error {
	_, existe := d[palavra]
	if !existe {
		return ErrPalavraInexistente
	}
	d[palavra] = definicao
	return nil
}

// Busca - realiza uma busca no dicionario pela palavra chave e retorna a definição da mesma
func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

// Adiciona - realiza uma adição de uma palavra no dicionario junto com sua definição
func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, existe := d[palavra]
	if existe {
		return ErrPalavraExistente
	}
	d[palavra] = definicao
	return nil
}
