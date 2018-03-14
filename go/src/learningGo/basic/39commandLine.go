package main

import (
	"os/exec"
	"fmt"
)

func main() {
	dateCmd := exec.Command("date")
	dateOuput, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	//fmt.Println(dateOuput) it will brint byte slice
	fmt.Println(string(dateOuput))
}
