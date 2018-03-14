package main

import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main() {
	h := sha1.New()
	s := "Hash this one..yoyo"
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("%x",bs)//hash is in usually hex string....so just hex
	//just like sha1 we have md5 hash
	m := md5.New()
	m.Write([]byte(s))
	bs = m.Sum(nil)
	fmt.Println()
	fmt.Printf("%x", bs)
}
