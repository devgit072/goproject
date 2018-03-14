package main

import "fmt"

func main() {
	fmt.Println(oneWayReturn())
	fmt.Println(multipleReturnValue())
	//how can u use these multiple return values
	a,b,c := multipleReturnValue()
	fmt.Println(a,b,c)
	//in some case we simply want to ignore one parameter and use only other one
	_,k,l := multipleReturnValue()
	fmt.Println(k,l)

	m,n := namedReturn()
	fmt.Println(m,n)
}

func oneWayReturn() int{
	return 12
}

func multipleReturnValue() (int, string, float32) {
	return 12,"hello", 78.9
}

func namedReturn() (a, b int) { //here we are declaring the variables and returning same
	a = 12
	b = 13
	return
}
