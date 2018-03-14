package main

import "fmt"

func main() {
	variadicFunc(2,3,11,34,1,11,34,11)
	//we can even pass slices/array in variadic function
	slice1 := []int {1,9,23,1,76,33,223,677}
	variadicFunc(slice1...)//note the three dots after slice variable name
}

func variadicFunc(params...int)  {
	for _,ele := range params{
		fmt.Println(ele)
	}
}
