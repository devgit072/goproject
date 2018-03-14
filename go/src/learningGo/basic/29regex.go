package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main() {
	match, _ :=regexp.MatchString("p([a-z]+)ch","peach1")
	fmt.Println(match)
	r,_ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	in :=[]byte("a peach")
	out :=r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))
}
