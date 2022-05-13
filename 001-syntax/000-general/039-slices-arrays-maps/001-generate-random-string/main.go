package main

import (
    "fmt"
    "math/rand"
    "time"

)

func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randInt(97, 122))
    }
    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func main() {

    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println(randomString(12))
}
