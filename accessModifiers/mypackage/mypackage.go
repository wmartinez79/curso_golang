package mypackage

import "fmt"

// Car struct with public access
type PublicCar struct {
	Brand string
	Year  int
	Model string
}

// car struct with private access
type privateCar struct {
	brand           string
	year            int
	model           string
	calculatedValue float64
}

var car privateCar

// set value of private car
func SetPrivateCar(brand string, year int, model string) {
	car.brand = brand
	car.year = year
	car.model = model
	calculateCarValue()
}

func calculateCarValue() {
	if car.year < 2020 {
		car.calculatedValue = 20000
	} else {
		car.calculatedValue = 50000
	}
}

// print value of private car
func PrintPrivateCarValue() {
	fmt.Println("current value of private car ", car)
}
