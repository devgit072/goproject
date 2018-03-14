package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println // this is way to create alias in go...a funky way

func main() {
	p(strings.ToLower("HellOJKJK"))
	p(strings.Contains("hello","ll"))
	p(strings.ToTitle("hello dear"))// it means toUpper
	p(strings.Split("hello dera whats up hjj k jj     jj", " "))
}


