package main

import(
	"fmt"
)

func main() {
	var scores [10]int
	scores[0] = 42
	fmt.Println("scores: ", scores)

	scores2 := [4]int{9001, 42, 13, 27}
	for index, value := range scores2 {
		fmt.Printf("index: %d - value: %d\n", index, value)
	}
}