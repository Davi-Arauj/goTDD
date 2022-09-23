package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "Mundo")
}

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func main() {
	if err := http.ListenAndServe(":5001", http.HandlerFunc(HandlerMeuCumprimento)); err != nil {
		fmt.Println(err)
	}
}
