package main

import "fmt"

func main() {
	//just declaration of map can be done by built in make, here no size specified
	map1 := make(map[string]int)
	map1["a1"]=12
	map1["b1"]=34
	map1["b3"]=34
	fmt.Println(map1)

	//here we can declare and initialize map in same line
	var map2 = map[int]int {12:11,90:21,6:12}
	map2[1]=11
	fmt.Println(map2)
	fmt.Println(map2[90])//elements can be get as map[key]
	fmt.Println(len(map2))

	//built in delete deletes the element from map
	delete(map2,12)
	delete(map1,"a1")
	fmt.Println(map2)
	fmt.Println(map1)

	val := map2[90]
	fmt.Println(val)
	//if the map doesnt contain key, then value is retured as 0 or false depending upon value type of map.
	val1 := map2[88]
	fmt.Println(val1)
}
