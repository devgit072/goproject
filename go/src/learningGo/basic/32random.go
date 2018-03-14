package main

import (
	"fmt"
	"math/rand"
)

//rand is lib in golang which uses random number generation...
func main() {
	fmt.Println(rand.Intn(100))// from 0 to 100
	fmt.Println(rand.Int())
	fmt.Println(rand.Float32())// from 0.0 to 1.0
	fmt.Println((rand.Float32()*5)+5)//from 5.0 to 10.0
	fmt.Println()
}