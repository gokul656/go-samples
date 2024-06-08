package main

import (
	"fmt"
	"time"

	"github.com/gokul656/go-samples/logger"
)

func init() {
	logger.Log("intialzing channels")
}

// a channel is always blocking if it's full
// it's always good to close channels from the sender side
// number of readers should be equal to number of writers

func UnBufferedChannels() {
	c := make(chan struct{}) // open

	// sender
	go func() {
		defer close(c)
		c <- struct{}{} // send
		fmt.Println("Sent first value")

		c <- struct{}{}
		fmt.Println("Sent second value")
	}()

	for data := range c { // blocking
		fmt.Println(data)
	}
}

func BufferedChannel() {
	c := make(chan struct{}, 3)
	c <- struct{}{}
	c <- struct{}{}

	c <- struct{}{}

	one := <-c // block
	two := <-c // block

	three := <-c

	fmt.Printf("%s %s %s", one, two, three)
}

func BufferedChannelExample() {
	// Create a buffered channel with a capacity of 5 (rate limit)
	requests := make(chan int, 5)

	go func() {
		for i := 0; i < 1; i++ {
			request := <-requests
			fmt.Println("Processing request", request)
			time.Sleep(200 * time.Millisecond) // Simulate processing time
		}
	}()

	// Simulate incoming requests
	for i := 1; i <= 10; i++ {
		select {
		case requests <- i:
			// Request accepted
			fmt.Println("Request", i, "accepted")
		default:
			// Channel is full, rate limit reached
			fmt.Println("Rate limit reached, request", i, "blocked")
			time.Sleep(200 * time.Millisecond) // Simulate processing time
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Consume requests from the channel
	for i := 0; i < cap(requests); i++ {
		request := <-requests
		fmt.Println("Processing request", request)
		time.Sleep(200 * time.Millisecond) // Simulate processing time
	}
}

func SelectExample() {

	market := make(chan struct{})
	ticker := make(chan struct{})

	go func() {
		for {
			select {
			case m := <-market:
				fmt.Printf("incoming msg from market %s\n", m)
			case t := <-ticker:
				fmt.Printf("incoming msg from ticker %s\n", t)
			}
		}
	}()

	producer(market, ticker)

	close(market)
	close(ticker)
}

func producer(market, ticker chan struct{}) {
	for i := range 10 {
		if i%2 == 0 {
			market <- struct{}{}
		} else {
			ticker <- struct{}{}
		}
	}
}
