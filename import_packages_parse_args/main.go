package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("First argument", os.Args[0])
	
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}