package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("yo")
	file := createFile("./test/defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("creating file")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("writing file")
	fmt.Fprintln(file, "hello and welcome, this is a test file..")
}

func closeFile(file *os.File) {
	fmt.Println("closing file")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
