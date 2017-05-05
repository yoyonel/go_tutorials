package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

// Structures dont have constructors.
// Instead, you create a function that 
// returns an instance of the desired type (like factory):
func PtrNewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name: name,
		Power: power,
	}
}

func main() {
	ptr_gohan := PtrNewSaiyan("Gohan", 1000)
	fmt.Println(*ptr_gohan)

	gohan := NewSaiyan("Gohan", 1000)
	fmt.Println(gohan)
}