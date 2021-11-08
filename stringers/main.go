package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("I have %d GB RAM, %d GB hard drive and it is a %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 500, brand: "Lenovo"}
	fmt.Println(myPc)

}
