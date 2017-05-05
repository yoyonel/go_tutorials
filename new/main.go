package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

func main() {
	declare_new_variable_pointer_1()
	declare_new_variable_pointer_2()
}

func declare_new_variable_pointer_1() {
	// Short version 
	goku := &Saiyan{
		Name: "goku",
		Power: 9000,
	}
	fmt.Println(*goku)
}

func declare_new_variable_pointer_2() {
	// Short version 
	goku := new(Saiyan)
	goku.Name = "goku"
	goku.Power = 9001

	fmt.Println(*goku)
}