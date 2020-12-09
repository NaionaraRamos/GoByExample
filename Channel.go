package main

import (
	"fmt"
	"time"
)

//Channel Synchronization
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
//Channel Synchronization

//Channel Directions
func ping(pings chan<-string, msg string) {
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}
//Channel Directions

func main() {

	//messages := make(chan string)

	//Channel Buffering
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//Channel Buffering

	//go func() { messages <- "ping"}()

//	msg := <- messages
//	fmt.Println(msg)


	//Channel Synchronization
	done := make(chan bool, 1)
	go worker(done)

	<- done
	//Channel Synchronization


	//Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
	//Channel Directions
}