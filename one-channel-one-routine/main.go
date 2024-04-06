package main

import (
	"fmt"
	"time"
)

func treatsServer(bowl chan string, numTreats int) {
	fmt.Printf("treatsServer: Preparing %d treats. \n\n", numTreats)
	time.Sleep(1 * time.Second)
	// Serve numTreats treats
	for i := 1; i <= numTreats; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("treatsServer: Treat number %d served now.\n", i)
		bowl <- fmt.Sprint("treat")
	}
}

func main() {
	bowl := make(chan string)
	numTreats := 5
	fmt.Printf("main: Hey, it's your Doggo, give me %d treats.\n\n", numTreats)

	// Serve treats
	go treatsServer(bowl, numTreats)

	// Wait for treats
	for j := 1; j <= numTreats; j++ {
		fmt.Printf("main: I got %d %ss. Woof!\n\n", j, <-bowl)
	}

	fmt.Println("main: Bark later, alligator!")
}
