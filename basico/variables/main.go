package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi 1:", pi)
	fmt.Println("Pi 2:", pi2)
	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero Values los valoreres que toman valores por defecto en go no son null o undefined
	// a = 0
	// b = 0
	// c = ''
	// d = false
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
