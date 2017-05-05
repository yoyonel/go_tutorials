package main

import (
	"fmt"
)

func main() {
	use_func_return_values_1()
	use_func_return_values_2()
}

func use_func_return_values_1() {
	value, exists := power("goku")	
	if exists == false {
		// handle this error case
	}
	fmt.Printf("%d %d\n", value, exists)
}

func use_func_return_values_2() {
	_, exists := power("goku")
	if exists == false {
		// handle this error case
	}	
	fmt.Printf("%d\n", exists)
}

func log(message string) {

}

func add(a int, b int) int {
	return a + b
}

func power(name string) (int, bool) {
	power := 9000
	exist := true
	return power, exist
}