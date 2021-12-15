package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

)

func read(f string) {
	verticalDir := 0
	horizontalDir := 0

	file, err := os.Open(f)
	if err != nil {
		log.Fatal("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		linePart := strings.Split(scanner.Text(), " ")
		dir := linePart[0]
		magnitude, err := strconv.Atoi(linePart[1])
		if err != nil {
			log.Fatal("Failed parsing the magnitude: %s", err)
		}

		switch dir {
		case "up":
			verticalDir -= magnitude
		case "down":
			verticalDir += magnitude
		case "forward":
			horizontalDir += magnitude
		default:
			fmt.Printf("invalid case %s - %d", dir, magnitude)
			continue
		}
	}
	fmt.Println(verticalDir, horizontalDir)
}

func main() {
	f := "file.txt"
	read(f)
}
