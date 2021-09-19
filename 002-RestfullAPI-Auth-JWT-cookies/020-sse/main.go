package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tmaxmax/go-sse"
)

func server() {
	s := sse.NewServer()

	go func() {
		m := sse.Message{}
		m.AppendText("hi everyone !")

		for range time.Tick(time.Second) {
			_ = s.Publish(&m)
		}
	}()

	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	server()
}
