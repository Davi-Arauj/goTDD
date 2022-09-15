package arrayslices

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		total := Soma(numeros)
		esperado := 6

		if total != esperado {
			t.Errorf("esperado '%d', mas o resultado foi '%d'", esperado, total)
		}
	})
}

func ExampleSoma() {
	numeros := []int{1, 1, 1, 1}
	total := Soma(numeros)
	fmt.Println(total)
	// Output: 4
}
