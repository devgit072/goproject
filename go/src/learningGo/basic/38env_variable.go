package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	//set env variable
	os.Setenv("dev","raj")
	fmt.Println(os.Getenv("dev"))
	fmt.Println(os.Getenv("dev1"))//no value set for the key

	//print the list of all env variables
	for _,e := range os.Environ() {
		pair := strings.Split(e,"=")
		fmt.Println(pair[0],"=",pair[1])
	}
}
