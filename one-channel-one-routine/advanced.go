package main

import (
	"fmt"
	"time"
)

const (
	treatPreparationTime = 1 * time.Second
	treatServingTime     = 1 * time.Second
)

// treatsServer prepares treats and serves them one by one until numTreats are served.
func treatsServer(bowl chan string, numTreats int) {
	// Prepare treats
	fmt.Printf("treatsServer: Preparing %d treats...\n\n", numTreats)
	time.Sleep(treatPreparationTime)

	// Serve treats
	for i := 1; i <= numTreats; i++ {
		time.Sleep(treatServingTime)
		fmt.Printf("treatsServer: Treat number %d served now.\n", i)
		bowl <- fmt.Sprint("treat")
	}
}

func main() {
	bowl := make(chan string)
	numTreats := 5
	fmt.Printf("main: Hey, it's your Doggo, give me %d treats.\n\n", numTreats)

	// Serve treats in a separate goroutine
	go treatsServer(bowl, numTreats)

	// Wait for treats
	for j := 1; j <= numTreats; j++ {
		fmt.Printf("main: I got %d %ss. Woof!\n\n", j, <-bowl)
	}

	fmt.Println("main: Bark later, alligator!")
}
