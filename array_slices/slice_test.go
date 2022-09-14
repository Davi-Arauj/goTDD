package arrayslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSomaTudo(t *testing.T) {

	verificarSoma := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado, '%v' esperado '%v'", resultado, esperado)
		}
	}
	t.Run("Soma tudo", func(t *testing.T) {
		resultado := SomaTudo([]int{1, 2, 3}, []int{3, 3, 3})
		esperado := []int{6, 9}
		verificarSoma(t, resultado, esperado)
	})

	t.Run("Soma todo o resto", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2, 3}, []int{3, 3, 3})
		esperado := []int{5, 6}
		verificarSoma(t, resultado, esperado)
	})

	t.Run("Soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 3, 3})
		esperado := []int{0, 6}
		verificarSoma(t, resultado, esperado)
	})
}

func ExampleSomaTudo() {
	resultado := SomaTudo([]int{1, 2, 3}, []int{3, 3, 3})
	fmt.Println(resultado)
	// Output: [6 9]
}

func ExampleSomaTodoOResto() {
	resultado := SomaTodoOResto([]int{1, 2, 3}, []int{3, 3, 3})
	fmt.Println(resultado)
	// Output: [5 6]
}
