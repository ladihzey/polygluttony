package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./assets/files.txt")
	checkError(err)
	io.WriteString(file, "This is a test file")
	checkError(err)
	fmt.Printf("Wrote to file %v\n", file.Name())
	defer file.Close()
	defer readFile("./assets/files.txt")
}

func readFile(path string) {
	file, err := os.ReadFile(path)
	checkError(err)
	fmt.Println("File content:", string(file))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
