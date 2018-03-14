package main

import "fmt"

//goroutine is lightwieght thread

func main()  {
	printSome("direct")

	// go word is used to start a new goroutine
	go printSome("goroutine") //here new execution path starts , so total path is now two and both will execute concurrently
	for i:=0;i<4;i++ {
		fmt.Println("In main")
	}
	//you can start anonymous function call with goroutine
	go func(str string) {
		fmt.Println("Its anonymous function call in goroutine")
	}("Anon")
	fmt.Scanln() // Scanl is just for wait enough to execute goroutine, else it might happen that main thread finishes early
	//and goroutine will not get enough time to execute, above if u give i<300 then goroutine will get enough time to execute

	fmt.Println("Done")
}

func printSome(str string)  {
	for i:=0;i<4;i++ {
		fmt.Println(str, i)
	}
}

