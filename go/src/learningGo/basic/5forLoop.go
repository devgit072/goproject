package main

import "fmt"

/*
for is only loop in golang unlike other language where we have loop constructs like for, while , do-while etc
 */

 /*
 We can use for loop in three ways:
  */
func main() {
	// one way
	var i = 1
	for i<10 {
		fmt.Println("YAYA",i)
		i++
	}

	fmt.Println()
	//second way
	for j:=0;j<10;j++ {
		fmt.Println("YOYO", j)
	}
	fmt.Println()

	k:=0
	for {
		fmt.Println("Keep looping...")
		k++
		if k > 3 {
			break
		}
	}
}
