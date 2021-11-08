package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100

	fmt.Println(a)

	myPc := pc{ram: 16, disk: 500, brand: "Lenovo"}
	fmt.Println(myPc)

	myPc.ping()

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)
}
