package main

const (
	prefixoOlaPortugues = "Olá, "
	prefixoOlaFrances   = "Bonjour, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoOlaRusso     = "Привет, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case "Frances":
		prefixo = prefixoOlaFrances
	case "Espanhol":
		prefixo = prefixoOlaEspanhol
	case "Russo":
		prefixo = prefixoOlaRusso
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
func main() {

}
