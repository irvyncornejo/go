package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//otra manera de instanciar
	var otherCar car
	otherCar.brand = "Porche"
	fmt.Println(otherCar)

}
