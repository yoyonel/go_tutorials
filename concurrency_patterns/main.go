// Google I/O 2012 - Go Concurrency Patterns 
// url: https://www.youtube.com/watch?v=f6kdp27TYZs
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// boring("boring!")

	go boring("boring!")
	fmt.Println("I'm listenning.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring. I'm leaving.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
