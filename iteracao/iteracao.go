package iteracao

// Repetir repeti o caracter recebido de acordo com a quantidade passada
func Repetir(caracter string, quantidade int) (repeticoes string) {
	for i := 0; i < quantidade; i++ {
		repeticoes += caracter
	}
	return repeticoes
}
