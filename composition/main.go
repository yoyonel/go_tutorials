package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, i'm %s\n", p.Name)
}

// Composite type 'Person' (without explicit field name)
type Saiyan struct {
	*Person
	Power int
}

// OverLoading (composite function Introduce from Person struct)
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func main() {
	// The Saiyan structure have a field type of *Person.
	goku := Saiyan{
		Person: &Person{"Goku"},
		Power: 9001,
	}
	// Because we didn't give it an explicit field name, we can implicitly access
	// the fields and functions of the composed type (Person).
	goku.Introduce()

	fmt.Printf("%s\n", goku)
	
	// However, the Go compiler did give it a field name, consider the perfectly valid:
	fmt.Printf("%s\n", goku.Name)
	fmt.Printf("%s\n", goku.Person.Name)

	goku.Person.Introduce()	
}