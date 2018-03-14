package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.RFC1123))

	t1, _ := time.Parse(time.RFC3339, "2012-11-08T19:08:53+00:00")
	fmt.Println(t1)
	//apart from RFC3339 we have other format as well for date&time representation
}
