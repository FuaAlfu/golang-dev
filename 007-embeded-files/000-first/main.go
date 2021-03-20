package main

import _ "embed"

func main() {
	//go:embed yo.txt
	var s string
	print(s)
}
