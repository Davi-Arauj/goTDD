package arrayslices

// SomaTudo recebe N slices e retorna a soma dos seus numeros em outro slice
func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

// SomaTodoOResto recebe N slices e retorna a soma dos seus numeros sem incluir o primeiro numero do slice
func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			somas = append(somas, Soma(numeros[1:]))
		}
	}
	return somas
}
