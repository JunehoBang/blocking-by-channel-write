package main

import (
	"log"
	"time"
)

var addListenerChan chan int

func worker() {
	for {
		time.Sleep(time.Second)
		received := <-addListenerChan
		log.Println("received: ", received)
	}

}

func main() {
	addListenerChan = make(chan int)
	// in case the length of the channel is less than 2,
	// the  the input code below (<-3) causes the main function
	// to be blocked.

	// One can check the "transfer complete" message is printed
	// after the "received" message is printed by the worker

	//To have the main thread and the worker thread to work independent
	//regardless of whether the channel has the value, the channel's length should
	//be set larger than or equal to 2
	//That is, [at line 20]: addListenerChan = make(chan int, 2)

	go worker()
	addListenerChan <- 3
	log.Println("transfer complete")
	time.Sleep(2 * time.Second)
	log.Println("done!")
}
