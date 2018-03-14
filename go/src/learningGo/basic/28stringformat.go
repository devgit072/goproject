package main

import (
	"fmt"
	"os"
)

type point struct {
	x,y int
}
func main() {
	p := point{12,99}

	fmt.Printf("struct value : %v", p) // v is used for struct
	fmt.Println()
	fmt.Printf("struct value with field name : %+v", p)// field name for struct
	fmt.Println()
	fmt.Printf("source code is : %#v\n", p) // source code
	fmt.Println()
	fmt.Printf("Type of variable is : %T\n", p) // type of variable
	fmt.Printf("boolean is : %t\n", true) // boolean formatting
	fmt.Printf("hello here integer value : %d\n", 123)
	fmt.Printf("binary value of 14 is %b \n", 20)
	fmt.Printf("character corresponding to given integer: %c \n", 39)
	fmt.Printf("Hex encoding : %x \n", 29)
	fmt.Printf("Float representaion: %f \n", 34.9)
	fmt.Printf("float in e format %e\n", 88787.888838824)
	fmt.Printf("float in E format %E\n", 556662.12234432)
	fmt.Fprintf(os.Stderr, "Error is %s", "hello")

}
