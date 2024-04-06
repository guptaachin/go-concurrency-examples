package main

import (
	"fmt"
	"time"
)

func serve_food(bowl chan string) {
	for i := 0; ; i++ {
		fmt.Println("Serving food in 1 sec...")
		time.Sleep(1 * time.Second)
		bowl <- fmt.Sprint("treat")

	}
}

func main() {
	bowl := make(chan string)
	fmt.Println("Hey, give me some treats.")

	// Serve food
	go serve_food(bowl)

	// Consume food
	for j := 0; j < 5; j++ {
		fmt.Printf("I got %d %ss. Woof!\n", j+1, <-bowl)
	}

	fmt.Println("Bark later, alligator!")
}
