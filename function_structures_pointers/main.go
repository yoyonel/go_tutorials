package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

func main() {
	pass_argument_by_copy()
	pass_argument_by_pointer()
	pass_argument_by_pointer_2()
}

func pass_argument_by_copy() {
	goku := Saiyan{"Goku", 9003}
	Super_by_copy(goku)
	fmt.Println(goku)	
}

func Super_by_copy(s Saiyan) {
	s.Power += 10000
}

func pass_argument_by_pointer() {
	// goku := &Saiyan{"Goku", 9004}
	// Super_by_pointer(goku)
	
	goku := Saiyan{"Goku", 9004}
	Super_by_pointer(&goku)

	fmt.Println(goku)	
}

func Super_by_pointer(s *Saiyan) {
	s.Power += 10000
}

func pass_argument_by_pointer_2() {
	goku := Saiyan{"Goku", 9005}

	Super_by_pointer_2(&goku)
	// print out: {Goku 9005}
	fmt.Println(goku)

	Super_by_pointer_3(&goku)
	// print out: {Goku 9005}
	fmt.Println(goku)	
}

func Super_by_pointer_2(s *Saiyan) {
	// Assign to pointer variable 's' a new address to new struct values
	// but we don't change the values of the original address pointer (from the call)
	s = &Saiyan{"Gohan", 1000}
}

func Super_by_pointer_3(s *Saiyan) {
	// Assign a new value to the address contain in pointer variable 's' (from call)
	*s = Saiyan{"Gohan", 1000}
}
