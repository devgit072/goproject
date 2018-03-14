package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("hello1")
	fmt.Println("hello2")
	fmt.Println("hello3")

	done <- true
}

func main() {
	channel_sync := make(chan bool, 1)
	worker(channel_sync)
	<-channel_sync// release the message , then only it will be un-blocked......
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
	fmt.Println("This will start only above is finished...")
}
