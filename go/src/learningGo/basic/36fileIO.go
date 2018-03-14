package main

import (
	"os"
	"math/rand"
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
)

func main() {
	//create one file
	file_name := fmt.Sprintf("/Users/devrajsingh/testDir/%s.txt", rand.Intn(10000))
	os.Create(file_name)
	readFile()
	writeFile()
	readFromStdIn()
}

func readFile() {
	//read one existing file
	data, err := ioutil.ReadFile("/Users/devrajsingh/testDir/testFile.txt")
	checkErr(err)
	//when u read from file it reads bytes..
	//so data has bytes.....so you need to convert bytes into string
	fmt.Println(string(data))

	fyle,err :=os.Open("/Users/devrajsingh/testDir/testFile.txt")
	checkErr(err)
	b1 := make([]byte, 5)//make byte array...5 bytes means five character will be read
	n1, err := fyle.Read(b1)//in b1 data is read
	fmt.Printf("%d bytes and data is %s", n1, string(b1))
	fyle.Seek(6,0) // seek is point file pointer to 6th place
	checkErr(err)
}

func writeFile() {
	//byte will be read as data from file and byte will be written as data in file
	//so comvert string into byte and write that in file
	b := []byte("Bangalore real estate sucks....")
	err := ioutil.WriteFile("/Users/devrajsingh/testDir/testFile1.txt", b, 0644)
	checkErr(err)

	f, err := os.Create("/Users/devrajsingh/testDir/testFile2.txt")
	checkErr(err)
	f.WriteString("At least there will be code...")
}

func readFromStdIn() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		inputFromUser := scanner.Text();
		fmt.Println(strings.ToUpper(inputFromUser))
		i++
		if i > 3 {
			break
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}