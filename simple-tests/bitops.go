package main

import "fmt"

type Prueba struct {
	name 	string
	age 	uint8
}

func (p Prueba) sayhello()  {
	fmt.Printf("Hello, %s\n", p.name)
}

func RunTest(i uint8)  {

	fmt.Printf("%08b\n", i)
	fmt.Printf("%08b\n", i << 3)
	fmt.Printf("%08b\n", i >> 3)

	fmt.Print("\U0001f475\n")
	fmt.Print("🦠\n")

}

func main() {
	RunTest(0xFF)

	var p = Prueba{"eva", 7}
	p.sayhello()
}
