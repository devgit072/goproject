package main

import (
	"time"
	"fmt"
)

//tickers are used when we do things repeatedly at regular intervals
func main() {
	tickers := time.NewTicker(500 * time.Millisecond) // this will tick every 500 millisecond
	//starts a new routine which will use tickers to print stuff at every 500 millisecond
	go func() {
		for t := range tickers.C { // this is way to use tickers
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	tickers.Stop() // when ticker stops, ticker exit and if u dont use this line and then ticker might print one more time since
	//waiting is 2000 ms and tickers tick every 500 ms
	fmt.Println("Done and ticker stopped")
}


