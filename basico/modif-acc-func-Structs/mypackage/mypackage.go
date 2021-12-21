package mypackage

import "fmt"

//Public car con acceso publico
type Car struct {
	Brand string
	Year  int
}

//Función pública
func Hola(brand string) {
	var myCar car
	myCar.brand = brand
	myCar.year = 1990
	fmt.Println("Función pública desde el paquete")
	fmt.Println("Valor desde la estructura privada")
	privadaHola(myCar)
}
func privadaHola(p car) {
	fmt.Println("Función privada desde el paquete")
	fmt.Println(p)
}

//Struct privado que en esta parte no se podra acceder desde el main
type car struct {
	brand string
	year  int
}
