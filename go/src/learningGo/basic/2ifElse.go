package main

import "fmt"

func main() {

	var a = 10
	if a%2==0 {
		fmt.Println("A is even....")
	} else {
		fmt.Println("A is odd....")
	}

	if k:=45;k>100 { // this is one way where u can declare and initialize variable in if block and use it
		fmt.Println(k, " is greater than 45")
	} else {
		fmt.Println(k, "is lesser indeed")
	}

	var j=90
	if j<90 {
		fmt.Println("YOYO")
	} else if j==90 {
		fmt.Println(j," is equal to 90")
	} else {
		fmt.Println("KOKO")
	}

	//remember, there is no ternary in go
}
