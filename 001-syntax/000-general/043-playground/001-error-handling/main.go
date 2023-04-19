package main

import(
	"fmt"
	"log"
)

func doSomething() error {
    // some code that might return an error
    if err != nil {
        return fmt.Errorf("an error occurred: %v", err)
    }
    return nil
}

func main() {
    if err := doSomething(); err != nil {
        // handle the error
        log.Fatalf("error occurred: %v", err)
    }
    // continue with program execution
}