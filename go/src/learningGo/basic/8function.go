package main

import "fmt"

func main() {
	oneway(12,"hello",8.9)
	secondWay(9,12,89)
	thirdway(12,34,56,"hello")
}

func oneway(a int, b string, f float32) {//functions can be unused unlike variables
	fmt.Println(a,b,f)
}

func secondWay(a, b, c int) {
	fmt.Println(a,b,c)
}

func thirdway(a, b, c int, str string) {
	fmt.Println(a,b,c,str)
}
