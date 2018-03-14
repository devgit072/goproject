package main

import "fmt"

func main() {
	array := []int {2,5,1,4,6,7}
	//so built in 'range' over array gives two things one index and other element at that index
	for index,ele := range array {
		fmt.Println(index, ele)
	}
	fmt.Println("\n\n")
	for _, ele := range array { // If you don't want index only element than you can leave index as optional as _
		fmt.Println(ele)
	}
	//in case if you want only index than
	for index := range array {
		fmt.Println(index)
	}
	//OR
	for index,_ := range array {
		fmt.Println(index)
	}

	// range over maps can be done like this
	map1 := map[string]string {"a":"b", "c": "d", "e":"f", "j":"n"}
	for k,v := range map1 {
		fmt.Println(k,v)
	}

	//If you want to print take only key then :
	for k := range map1 {
		fmt.Println(k)
	}
}