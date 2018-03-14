package main

import (
	"os"
	"fmt"
)

//panic in go is like System.exit(0) in Java....its used when u don't know how to handle the error and simply wants to
//quit...panic is same: print same and exit with non-zero status

func main() {
	fmt.Println("This will be executed....")
	openFile()
	fmt.Println("This will not be printed....bacause panic is executed before that")
}

func openFile() {
	_,err := os.Create("/tmp/myfile/lolo")//it will not create this file, so file IO exception will come
	if err != nil {
		panic("File IO came....I'm exiting")// it will be printed with Process finished with exit code 2 i.e status_code=2
	}
}
