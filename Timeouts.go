package main

import (
	"fmt"
	//"time"
)

func main() {

	//Timeouts

	/* c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	} */

	//Timeouts


	//Non-Blocking Channel Operations
/* 	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")	
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")	
	}

	
	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")	
	} */
	//Non-Blocking Channel Operations


	//Closing Channels

/* 	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<- done */


	//Range over Channels

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for e := range queue {
		fmt.Println(e)
	}
	//Range over Channels

	//Closing Channels
}