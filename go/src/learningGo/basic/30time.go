package main

import (
	"fmt"
	"time"
)
// time.Now() has all the information..
func main() {
	p := fmt.Println
	t := time.Now()
	p(t)
	fmt.Println(t.Second())
	fmt.Println(t.Minute())
	fmt.Println(t.Hour())
	fmt.Println(t.Weekday(), t.Month(),t.Day(), t.Year(), t.YearDay()/*which day in 365days sequence like 67th day in this year*/)
    oneTime := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    fmt.Println(oneTime)
    fmt.Println(oneTime.After(t))// it will simply print if oneTime is after now or not..a boolean question
    fmt.Println(oneTime.Sub(t)) // it says duration between two times
    diff := oneTime.Sub(t)
    fmt.Println(oneTime.Add(diff))
    //to subtract to have to do add(-diff)
    fmt.Println(oneTime.Add(-diff))

    //epoch time
    fmt.Println(time.Now().Unix())
    fmt.Println(time.UnixDate)
    fmt.Println(time.Now().UnixNano())
    fmt.Println(time.Unix(100,200))// it means add 100 seconds and 200 nanoseconds in 1 Jan, 1970
}
