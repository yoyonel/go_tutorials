package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

func main() {
	
	gohan := &Saiyan {
		Name: "Gohan",
		Power: 1000,
		Father: &Saiyan {
			Name: "Goku",
			Power: 9001,
			Father: nil,
		},
	}

	// url: http://stackoverflow.com/questions/4938612/how-do-i-print-the-pointer-value-of-a-go-object-what-does-the-pointer-value-mea
	fmt.Printf("%p - %v\n", gohan, *gohan)
	fmt.Printf("%p - %v\n", gohan.Father, *gohan.Father)
}
