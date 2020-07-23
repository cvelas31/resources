package main

import (
	"fmt"
	"math/rand"
)

func testUint32() {

}

func main() {
	fmt.Println("simpleRngRand")
	var num uint32 = 42
	for i := 0; i < 20; i++ {
		num = simpleRngRand(num)
		fmt.Println(num)
	}
	fmt.Println("")
	fmt.Println("NewPCG32: ")
	pcg := NewPCG32().Seed(3444837047, 2669555309)
	for n := 0; n < 10; n++ {
		fmt.Println(pcg.Random())
	}
	fmt.Println("")
	fmt.Println("mt19937_32go: ")
	Seed(42)
	for n := 0; n < 10; n++ {
		z := Float64()
		fmt.Println(z)
	}
	fmt.Println("")
	fmt.Println("Default math/rand Go: ")
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		aux := r1.Float64()
		fmt.Println(aux)
	}
}
