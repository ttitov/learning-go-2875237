package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Some text that has wrote to the file"
	file, err := os.Create("./my_file.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote to a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./my_file.txt")

}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println(string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
