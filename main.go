package main

import (
	"fmt"
	"time"
)

type KeyValuePair struct {
	Key   int
	Value float64
}

func main() {

	queue := make(chan KeyValuePair)

	go readQueue(queue)

	pushToQueue(queue)

	select {}
}

func readQueue(queue chan KeyValuePair) {
	// create a ticker with a 5-second interval
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		// use the ticker channel to block until the next tick
		<-ticker.C

		// try to receive a key-value pair from the queue
		select {
		case pair, ok := <-queue:
			if ok {
				fmt.Printf("Time: %s, Key: %d, Value: %f\n", time.Now().String(), pair.Key, pair.Value)
			} else {
				fmt.Println("Queue closed")
				return
			}
		default:
			fmt.Println("Waiting for message")
		}
	}
}

func pushToQueue(queue chan KeyValuePair) {
	// push some key-value pairs onto the queue
	queue <- KeyValuePair{Key: 500, Value: 0.015}
	queue <- KeyValuePair{Key: 200, Value: 0.05}

	// close the queue after the key-value pairs have been pushed
	//	close(queue)
}
