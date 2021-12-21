package main

import (
	"fmt"
	"strings"
)

func esPalindromo(text string) {
	var textRevers string
	textLower := strings.ToLower(text)
	for i := len(textLower) - 1; i >= 0; i-- {
		textRevers += string(textLower[i])
	}
	if textRevers == textLower {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}

}

func main() {
	slice := []string{"Hola", "Mundo", "go"}

	for _, valor := range slice {
		//fmt.Println(i, valor) si no queremos el indice lo comentamos
		fmt.Println(valor)
	}

	//Ejercicio palindromo
	esPalindromo("roma")
}
