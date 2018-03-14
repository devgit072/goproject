package main

import "fmt"

func main() {
	fmt.Println(person{12,"devraj", 5.6}) //struct has default way to initialize values (person{12,"devraj", 5.6}, note the parenthesis
	s1 := person{name:"devraj1", age:20, height:6.3} //this is another way to initialize the values and field can be in any order
	s2 := person{name:"devraj2"} // initialize only one field, thats okay other will have default value when u print the struct
	fmt.Println(s1)
	fmt.Println(s2)

	sp := &s1 //its pointer to struct, sp is pointer to struct
	fmt.Println(sp.name)
	fmt.Println(sp.age)
	// u can modify the value thru pointer
	sp.height=9.1 //structs are mutable
	fmt.Println(s1)
	fmt.Println(*sp)

	accessStructThruPointer(&s1)
	accessStruct(s1)
}

//struct is like class in golang

type person struct {
	age int
	name string
	height float32
}

//function can accept struct pointer
func accessStructThruPointer(ptr *person)  {
	fmt.Println(*ptr)
	ptr.height=8.9
	fmt.Println(*ptr)
}

func accessStruct(st person) {//its struct variable
	st.height=7.8
	fmt.Println(st)
}