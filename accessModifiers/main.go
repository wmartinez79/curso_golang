package main

import (
	pk "curso_go_platzi/accessModifiers/mypackage"
	"fmt"
)

func main() {
	var myCar pk.PublicCar
	myCar.Brand = "Hyudai"
	myCar.Model = "Grand i10"
	myCar.Year = 2020
	fmt.Println(myCar)
	pk.SetPrivateCar("Honda", 2022, "CRV")
	pk.PrintPrivateCarValue()
}
