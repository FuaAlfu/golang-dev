package main

import (
	"log"

	"github.com/gomodule/redigo/redis"

)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err) // handle error
	}
	defer c.Close()

	//fetching the data from redis cache
	// title, err := redis.String(conn.Do("HGET", "book:1", "title"))
	// if err != nil {
	// 	log.Fatal(err) // handle error
	// }
	// fmt.Println("Book Title:", title)
}
