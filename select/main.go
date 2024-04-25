package main

import (
	"fmt"
	"time"
)

// treatsServer simulates a treats provider that serves treats to a channel.
func treatsServer(bowl chan string) {
	// Infinite loop to continuously serve treats
	for i := 0; ; i++ {
		fmt.Printf("treatsServer: Serving treat now.\n\n")
		// Simulate serving a treat with a delay of 5 seconds
		time.Sleep(5 * time.Second)
		// Send a treat to the bowl channel
		bowl <- "treat"
	}
}

func main() {
	// Create a channel for receiving treats
	bowl := make(chan string)

	fmt.Printf("main: Give me a treat before going for a walk.\n\n")

	// Start the treats server goroutine
	go treatsServer(bowl)

	// Set a timeout of 3 seconds
	timeoutEating := time.After(3 * time.Second)

	// Use select statement to wait for either a treat or timeout
	select {
	case treat := <-bowl:
		// If a treat is received within the timeout
		fmt.Printf("main: I got a %s. Woof!\n\n", treat)
	case <-timeoutEating:
		// If timeout occurs before receiving a treat
		fmt.Println("main: Too long, let's go for a walk!")
	}
}
