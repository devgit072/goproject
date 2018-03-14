package main

import (
	"sort"
	"fmt"
)

func main() {
	//go has built in sort methods...
	arr1 := []string {"hell", "jkkk", "gheh", "dev", "hev", "dev"}
	sort.Strings(arr1)
	arr2 := []int {23,9,12,7,134,89,0,-89,-887,899}
	sort.Ints(arr2)
	fmt.Println(arr1)
	fmt.Println(arr2)
	//check if array are already in sorted order
	fmt.Println(sort.IntsAreSorted(arr2))
	fmt.Println(sort.StringsAreSorted(arr1))

}
