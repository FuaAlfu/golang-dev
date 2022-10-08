package main

import (
    "fmt"

)

func reverse(s []int) {
	start := 0
	end := len(s) - 1
    for start < end {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}



func main() {
  s := []int{892,111,7,203,42,989,812,75} 
  fmt.Printf("non reversed slice: %d\n", s)
  reverse(s)
  fmt.Printf("reversed slice: %d\n", s)
}