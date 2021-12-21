package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es uno")
	} else {
		fmt.Println("No es uno")
	}

	//with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es Verdad ambos son el valor")
	} else {
		fmt.Println("No son valor")
	}

	//with or
	if valor1 == 3 || valor2 == 5 {
		fmt.Println("Es Verdad uno de los valores es igual")
	} else {
		fmt.Println("Ningún valor es el de la condición")
	}
}
