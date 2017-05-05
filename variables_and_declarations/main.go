package main

import (
	"fmt"
)

func main() {
	simple_declaration()

	short_declaration()

	return_value()

	// error_double_declarations()

	assign_multiple_variables()

	no_error_double_declarations()

	// error_not_used_variable()
}

func simple_declaration() {
	// declare a integer variable 'power' (default value=0)
	var power int

	// assign a new value to the (existing) variable 'power'
	power = 9000

	// print out 'power' value
	fmt.Printf("Power=%d\n", power)
}

func short_declaration() {
	// Declare and assign value to a new variable 'power'
	power := 9001

	// print out 'power' value
	fmt.Printf("Power=%d\n", power)
}

func return_value() {
	power := getPower()

	// print out 'power' value
	fmt.Printf("Power=%d\n", power)}

func getPower() int {
	return 9002
}

// func error_double_declarations() {
// 	power := 9003
// 	// print out 'power' value
// 	fmt.Printf("Power=%d\n", power)

// 	// ERROR
// 	power := 9004
// 	// print out 'power' value
// 	fmt.Printf("Power=%d\n", power)	
// }

func assign_multiple_variables() {
	name, power := "Goku", 9004
	fmt.Printf("%s's power is over %d\n", name, power)
}

func no_error_double_declarations() {
	power := 9005
	fmt.Printf("The default power is %d\n", power)

	name, power := "Goku", 9006
	fmt.Printf("%s's power is over %d\n", name, power)	
}

// func error_not_used_variable() {
// 	name, power := "Goku", 9007
// 	// ERROR: no use of 'name' variable
// 	fmt.Printf("The default power is %d\n", power)
// }