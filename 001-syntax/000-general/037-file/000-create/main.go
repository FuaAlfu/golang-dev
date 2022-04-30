package main

import(
	"fmt"
	"os"
)

func main(){
	file := "./test.txt"

	createFile, err := os.Create(file)
	if err != nil{
		panic(err)
	}

	defer func() {
		createFile.Close()
	}()

	fmt.Fprintln(createFile,"testing..")
}