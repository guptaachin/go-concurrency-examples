package main

import (
	"fmt"
	"time"
)

func treatsServer(bowl chan string) {
	for i := 0; ; i++ {
		fmt.Printf("treatsServer: Serving treat now.\n\n")
		time.Sleep(5 * time.Second)
		// Serve treat
		bowl <- "treat"
	}
}

func main() {
	bowl := make(chan string)
	fmt.Printf("main: Give me a treat before going for walk.\n\n")
	// serve treats
	go treatsServer(bowl)

	timeoutEating := time.After(3 * time.Second)
	select {
	case treat := <-bowl:
		fmt.Printf("main: I got a %s. Woof!\n\n", treat)
	case <-timeoutEating:
		fmt.Println("main: Too long, let's go for a walk!")
	}
}
