package main

import (
	"fmt"
	"math"
)

func main() {
	values()
	fmt.Println()
	variables()
	constantType()
}

func values() {
//string concatenation
fmt.Println("Dev" + "raj")
fmt.Println(4+5)
fmt.Println(4.5/2.3)
fmt.Println(true || false)
fmt.Println(true && false)
}


func variables() {
	// Lets see numerous ways to see the declaration and initialization of variables
	var intType = 34  // declare and intialize
	var str = "Bengaluru" // declare and initialize
	var floatType float32  // first declare with its type and then initilize later in some other line
	floatType = 4.5 // initialize
	fmt.Println(intType, str, floatType)

	a:=34 // best way to declare and initialize. := can be used only for first time variable with initialization
	str1:= "Cohesity There"
	var str2 string = "Yayayayay" // redudant way for above stattements

	fmt.Println(a, str1, str2)
}

func constantType()  {
	// There are two types of literal in go: var and const, so either declare literal with var or const
	const a = 23 // now it can't be changed , and its not mandatory to use it unlike var
	fmt.Println(a)
	fmt.Println(math.Sin(34))
	const f = 345.0/13.6
	fmt.Println(f)
}
