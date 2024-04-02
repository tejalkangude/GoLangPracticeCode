package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This is New file and Writing file content in go lang files"

	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
