package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	n := 10
	repeticoes := Repetir("a", n)
	esperado := "aaaaaaaaaa"
	if len(repeticoes) != len(esperado) {
		t.Errorf("esperado '%d' mas obteve '%d'", len(esperado), len(repeticoes))
	}
	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func ExampleRepetir() {
	caracter := Repetir("b", 5)
	fmt.Println(caracter)
	// Output: bbbbb
}
