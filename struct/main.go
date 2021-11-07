package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Honda", year: 2013}
	fmt.Println(myCar)

	var secondCar car
	secondCar.brand = "Suzuki"
	fmt.Println(secondCar)
}
