package main

import (
	"encoding/json"
	"fmt"
	"time"

)

var t time.Time

func main() {
	t1 := time.Now()
	data, _ := json.Marshal(t1)
	json.Unmarshal(data, &t)
	fmt.Println(t == t1)
	fmt.Println(t1.Equal(t))
}
