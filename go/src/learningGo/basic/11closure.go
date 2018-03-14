package main

import "fmt"

func main()  {
	nextIntVar := nextInt() //It is variable type so its declared and initialized... i=0 and when u use this variable
	//again , anonymous function will be executed again and aagin whenever it will be used
	fmt.Println(nextIntVar()) // nextIntVar() ultimtely return int, so returned int will be printed
	fmt.Println(nextIntVar())
	fmt.Println(nextIntVar())

	nextIntVar1 := nextInt()
	nextIntVar2 := nextIntVar()
	fmt.Println(nextIntVar1())
	fmt.Println(nextIntVar2)
}
//closure is anonymous function
func nextInt() func() int { // here the return type is function which return int
	i:=0
	return func() int { //Its anonymous function inside a function, which operates on variable of parent function
		i++
		return i
	}
}