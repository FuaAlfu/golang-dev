package main

import "log"

func testFunc(s string){
	log.Println(s)
}

func main(){
  s := "testing.. ^"
  testFunc(s)

  //log
  log.SetFlags(log.Ldate | log.Lshortfile)
  log.Println(s)
  //log.Fatal("err")
 // log.Panic("panic.. ..")
}