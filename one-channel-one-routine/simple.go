package main

import (
	"fmt"
	"time"
)

// treatsServer prepares treat and serves it
func treatsServer(bowl chan string) {
	// Serve treats
	fmt.Printf("treatsServer: Preparing treat in 3 secs...\n")
	time.Sleep(3 * time.Second)
	fmt.Printf("treatsServer: Serving treat now.\n\n")
	time.Sleep(1 * time.Second)
	bowl <- fmt.Sprint("treat")
}

func main() {
	bowl := make(chan string)
	fmt.Printf("main: Hey, it's your Doggo, give me a treats.\n\n")

	// Serve treats in a separate goroutine
	go treatsServer(bowl)

	// Wait for treat
	fmt.Printf("main: I got a %s. Woof!\n\n", <-bowl)

	fmt.Println("main: Bark later, alligator!")
}
