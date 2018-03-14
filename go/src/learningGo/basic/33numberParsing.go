package main

import (
	"strconv"
	"fmt"
)

//strconv is used to parse string to number
func main() {
	f,_ :=strconv.ParseFloat("1.234",64)//float64
	f1,_ :=strconv.ParseFloat("1.234",32)//float32
	fmt.Println(f, f1)
	fmt.Println(strconv.ParseInt("68223", 0,64))// 0 means infer the base from string
	//from string if it is 0x23 then we can infer it is 16 base hex
	fmt.Println(strconv.ParseInt("0x1c8",0,32))
	n,e := strconv.ParseInt("677sgh243l", 0, 64)
	if(e != nil) {
		fmt.Println("error handled")
	} else {
		fmt.Println(n)
	}
	k,_ := strconv.Atoi("135")
	fmt.Println(k)
	l,er := strconv.Atoi("wat") //since parse unsuccefull then l will be 0 and er
	//will contain error message
	fmt.Println(er, l)
}