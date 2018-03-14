package main

import "fmt"

func main() {
	x := 5
	passByValue(x) // value will not change, since its is pass by value and the value is copied to new variable in called function
	fmt.Println(x)
	passByPointer(&x) // value will change since we are passing pointers.
	fmt.Println(x)

	// another way to use pointer is use "new" built in
	ptr := new(int) //it means ptr will act as pointer for int type variable, so type e.g int is must to define
	// new(int) is just used for declaration purpose but memory is allocated for that pointer
	assignValueToPointer(ptr)
	fmt.Println(*ptr)
}

func passByValue(x int) {
	x=0
}
func passByPointer(x *int) {
	*x=0 // it means that store value 0 in memory location of x, so original value will be modified
}

func assignValueToPointer(ptr *int)  {
	*ptr = 10 //it means 10 has been assigned to pointer ptr
}
