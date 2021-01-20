package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {

	// fmt.Println("hi there")
	if err := run(); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
}

func run() error {
	if 1 == 2 {
		return errors.New("just a random error")
	}

	return nil
}
