package main

import "fmt"

// channels are like pipe where concurrent goroutine uses to talk to each other
// think channel like message queue, where one goroutine will put message in queue and other will consume it

func main() {
	//how to make channel
	my_channel := make(chan string)

	//how to put message in channel
	go func() {my_channel <- "Here is message put in channel"}()

	//how to get message from channel to string var
	message_from_channel := <- my_channel // X->Y is sign of from X to Y
	// cant use, because message was already fetched out from channel once...message2 := <- my_channel

	fmt.Println(message_from_channel)
	channel_buffer()
}

func channel_buffer() {
	//by default channel are unbuffered, this means that only one item can be put in channel before it get consumed
	my_channel := make(chan string, 2)// by 2 we mean that we can put 2 message in channel
	my_channel <- "hello1"
	my_channel <- "hello2"

	fmt.Println(<-my_channel)
	fmt.Println(<-my_channel)
	//fmt.Println(<-my_channel) it will throw error because 3 message has been already consumed
}
