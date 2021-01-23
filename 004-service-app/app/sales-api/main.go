package main

import (
	"log"
	"os"
)

func main() {
	//prefix
	log := log.New(os.Stdout, "SALES : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	//adding an exported function
	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)

		//log.Fatalln(v ...interface{})
	}
}

func run(log *log.Logger) error {
	return nil
}
