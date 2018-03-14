package main

import (
	"fmt"
	"strings"
)

//lets create some function which will do some operation on collection

//first index of target string in given string slice
func indexOf(arr []string, t string) int{
	for i,s := range arr {
		if s==t {
			return i
		}
	}
	return -1
}

//returns true if any of the string matches target string i.e string is present in slice
func isPresent(arr []string, t string) bool {
	for _,s := range arr {
		if s==t {
			return true
		}
	}
	return false
}

// well what is predicate? Predicate is function which will accept input and return true or false
// predicate gives u answer in boolean whether given input satisfies condition mentioned in predicate function


// returns true if any of the string in slice satisfies the predicate
func Any(arr []string, predicate func(string) bool) bool{//notice how predicate written: <func_name> func(string) bool
	for _,s := range arr {
		if(predicate(s)) {
			return true
		}
	}
	return false
}

func All(arr []string, predicate func(string) bool) bool {
	for _,s := range arr{
		if !predicate(s) {
			return false
		}
	}
	return true
}

//filter all string which satisfies given predicate
func Filter(arr []string, predicate func(string) bool) []string {
	filtered_slice :=make([]string, 0)
	for _,s := range arr {
		if predicate(s) {
			filtered_slice = append(filtered_slice, s)
		}
	}
	return filtered_slice
}

// map function applies and modify the string in each and every string in given slice
func MapF(arr []string, modify func(string) string) []string {
	for i,s := range arr {
		arr[i] = modify(s)
	}
	return arr
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(indexOf(strs, "pear"))
	fmt.Println(isPresent(strs, "grape"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasSuffix(s, "p")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "p")
	}))
	fmt.Println(MapF(strs, strings.ToUpper))
}