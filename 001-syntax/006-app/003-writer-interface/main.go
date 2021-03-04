package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   using Writer
	   Writer interface{}
	*/
	//print all
	fmt.Println("hi")
	fmt.Fprintln(os.Stdout, "hi")
	io.WriteString(os.Stdout, "hi")

}
