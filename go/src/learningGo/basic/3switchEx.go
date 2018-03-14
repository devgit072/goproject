package main

import (
	"fmt"
	"time"
)

/*
Switch in go is alternative to if/else, so case can handle variable also apart from const

 */
func main() {
	i:=10

	switch i {
	case 20:
		fmt.Println(i,"its 10")
	case 10:
		fmt.Println(i,"is 10")
	case 90:
		fmt.Println("Well its 90")
	case 190,100: // in case u can give multiple values also
		fmt.Println("Done done")
	default:
		fmt.Println("None of above")
	}

	//here you will see , that case can handle variable also.
	//u can't use statement like, case i>10, but there is other way to use like in time, date etc
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Its friday")
	case time.Wednesday:
		fmt.Println("Well its Tuesday")
	case time.Saturday:
		fmt.Println("Saturday")
	default:
		fmt.Println("none")
	}

	// way how u can use variable statements in case. You switch block should be empty switch {case 10: } etc
	t:=time.Now()
	switch { //switch without any value
	case t.Hour() > 12:
		fmt.Println("Noon")
	default:
		fmt.Println("Morning")
	}

	k:=-100
	switch {
	case k > 10:
		fmt.Println("Hello world again")
	case k < 100:
		fmt.Println("KKKKK")
	case k < 90:
		fmt.Println("Lesser")
	default:
		fmt.Println("llllllll")

	}
}
