package main

import (
	"errors"
	"fmt"
)


//error is interface type
type errorArg struct {
	code int
	message string
}
//implement a function Error which can work on errorArg struct
func (e *errorArg) Error() string  {//Error is function of interface error
	return fmt.Sprintf("%d  - %s", e.code, e.message)
}

func main()  {
	fmt.Println(errorDemo(10))
	fmt.Println(errorDemo(20))
	fmt.Println(errorDemo(-20))

	// how to handle error
	for _,ele := range []int {10, -2,90,-9} {
		val,err := errorDemo1(ele)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(val)
		}
	}
}

func errorDemo(param int) (int, error) {
	if param < 0 {
		return -1, errors.New("Oops, its negative number") //its a kind of handling error
	} else {
		return param*10, nil
	}
}

func errorDemo1(param int) (int, error) {
	if param < 0 {
		return -1, &errorArg{code:10, message:"Negative number error"}
	} else {
		return param*10, nil
	}
}

