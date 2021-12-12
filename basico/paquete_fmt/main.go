package main

import "fmt"

func main() {
	//Declaracionde varibles
	helloMessage := "Hello"
	worldMessage := "World"
	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	//Si no se sabe el tipo de dato se pasa la variable con %v
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
}
