package arrayslices

// Soma soma os numeros passados e retorna o total
func Soma(numeros []int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}
