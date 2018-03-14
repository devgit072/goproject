package main

import (
	"os"
	"fmt"
)

func main() {

	params1 := os.Args // it will also consists of program name and parameters
	params2 := os.Args[1:]// it means from index 1 to remaining , that means it will
	//exclude the first parameter which is program name...we dont need that usually
	fmt.Println(params1)
	fmt.Println(params2)
    // first parameter usually looks something like that:
    // [/var/folders/g8/vpq4h7x941s_qhd7r1wfbxx40000gn/T/go-build249557170/command-line-arguments/_obj/exe/37commandLineArguments
}
