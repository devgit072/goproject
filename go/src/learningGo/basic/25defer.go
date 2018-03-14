package main

import (
	"fmt"
	"os"
)

//defer is like finally in java, defer <part_of_code> will be executed when program is finished and at last it will
// be executed, it main purpose is to resource reclaim, close file/conections etc

func main() {
	defer justPrint()// this is though first statement, but will be executed at last
	fmt.Println("Lets began")
	file := createFile("/Users/devrajsingh/testDir/testfile.txt")
	defer closeFile(file) // this will be executed after writeSomeThingInFile()
	writeSomeThingInFile(file)
}

func createFile(fileName string) *os.File{
	fmt.Println("Creating a new file....")
	f, err := os.Create(fileName)
	if err!=nil {
		panic("file not created")
	}
	fmt.Println("done with file creation...")
	return f
}
func writeSomeThingInFile(file *os.File) {
	fmt.Println("Writing something in file")
	fmt.Fprintln(file, "This is wrting on the wall..")
	fmt.Println("done with writing")
}
func closeFile(file *os.File) {
	fmt.Println("Closing the file....")
	file.Close()
}
func justPrint() {
	fmt.Println("Just printing for defer")
}
