package main

import ( 
"encoding/json"
    "fmt"

) 

var data = `{"name": "John", "age": 30, "city": "New York"}`

//Parse jsonData into a map  
var result map[string]interface{}

func main() { 
dataBytes := []byte(data)
json.Unmarshal(dataBytes, &result) 
name  := result["name"]
age   := result["age"]
city  := result["city"]
fmt.Println("Name: ", name) 
fmt.Println("City: ", city) 
fmt.Println("Age: ", age)
println("---")
for i, v := range result {
	fmt.Println("k: ,", i, " we are now adding ", v, "to value")
}
}
