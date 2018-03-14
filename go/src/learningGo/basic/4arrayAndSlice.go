package main

import (
	"fmt"
)

func main() {
	//what is diff between array and slices?
	//Both are same but with one diff... slice has no count mentioned and array has count mentioned while declaration

	var slices = []int {12,23,45,678}
	fmt.Println(slices)

	//both below are form of array
	var array = [5]int {12,23,34,121,3}//here we have mentioned fixed size of 5, more than that will throw error
	var varArray = [...]int {23,111,222,232,34343,2323,909,77}
	fmt.Println(array)
	fmt.Println(varArray)

	// built in function len() gives length of array
	fmt.Println("Length1:", len(slices))
	fmt.Println("Length2:", len(array))
	fmt.Println("Length3:", len(varArray))

	//just a declaration and no initialization
	var varSlice []int // if u don't mention size then its a 0 sized slice and varArr[2]=9 will throw error
	var varArr [3]int // by default all elments are initialized as 0 if u mentioned size, so printing this
	// array will print all elements as 0
	fmt.Println(varSlice)
	fmt.Println(varArr)
	varArr[0]=12
	varArr[1]=13
	fmt.Println(varArr)

	a:=[]int {1,2,3,4,5} //slice in one liner
	b:=[3]int {23,12,89} //array in one liner
	fmt.Println(a)
	fmt.Println(b)

	//two dimensional slice
	var twoDim [][]int
	fmt.Println(twoDim)
	/*
	for i:=0;i<5;i++ {
		for j := 0; j < 4; j++ {
			twoDim[i][j]=i*j
		}
	} this will throw array, because slice is something like immutable. So slices must be declaraed and initialized at same line
	*/
	var twoDimArray [3][2]int
	for i:=0;i<3;i++ {
		for j:=0;j<2;j++ {
			twoDimArray[i][j]=i*j
		}
	}
	fmt.Println(twoDimArray)
	differentWaysOfSlices()
}

func differentWaysOfSlices() {
	//slices can be created using built in default function:
	slice1 := make([]string, 0)// here slice has been initialized with 3 elements all are empty string
	//since slice is immuatble, so if you append new elements in slices append method will give u new instance of slice
	fmt.Println(slice1)
	slice1 = append(slice1, "hello")
	slice1 = append(slice1, "gello")
	slice1 = append(slice1, "kello")
	slice1 = append(slice1, "jjj")
	slice1 = append(slice1, "kkk")
	fmt.Println(slice1, len(slice1))//so it will length as 8: 3 empty string and 5 elements that we added later

	//copy of slice using built in function: copy
	copied_slice := make([]string, 8) // u need to be careful while giving size of copied slice,
	// if u give size 0 then 0 elements will be copied, if u give size a 4 then 4 elements will be copied
	copy(copied_slice, slice1)
	fmt.Println("The copied slice:", copied_slice)
	fmt.Println("The split slice:",slice1[0:5])// this means from 0th index including to 5th index excluding
	//as discussed earlier, we can d&i slice in same line
	d_i_slice := []string {"kkk", "kkkk", "aaaaa", "fhagsjhg"}
	fmt.Println(d_i_slice)

	twoDSlice := make([][]int, 5)
	for i:=0 ; i< 5 ;i++ {
		innerLen := i+1
		twoDSlice[i] = make([]int, innerLen);
		for j:=0 ;j<innerLen;j++ {
			twoDSlice[i][j]=i*j
		}
	}
	fmt.Println("Two d: \n", twoDSlice)
}
