package main

import (
	"fmt"

	"basico/mypackage/mypackage"
)

func main() {
	var myCar mypackage.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 1992
	fmt.Println(myCar)
	mypackage.Hola("Ducati")
}
