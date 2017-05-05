package main

import(
	"fmt"
)

type Saiyan struct {
	Name string
	Friends map[string]*Saiyan
}

func main() {
	map_1()
	map_2()
	map_with_structures()
	map_declare_initialize()
	map_iterate()
}

func map_1() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001

	fmt.Println("lookup map:", lookup)

	power, exists := lookup["vegeta"]

	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println("Search 'vegeta' key in lookup:",power, exists)

	// to get the number of keys
	total := len(lookup)
	fmt.Println("Number of keys in lookup map:", total)

	// to remove a value based on its key
	fmt.Println("delete goku key ...")
	delete(lookup, "goku")

	fmt.Println("lookup map:", lookup)
}

func map_2() {
	// Map grow dynamically.
	// However, we can supply a second argument to make to set an initial size
	lookup := make(map[string]int, 100)
	// defining initial size can help with performance.
	fmt.Println("lookup map:", lookup)	
}

func map_with_structures() {
	goku := &Saiyan{
		Name: "goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{"krillin", nil}

	fmt.Println(*goku)
}

func map_declare_initialize() {
	lookup := map[string]int{
		"goku": 9001,
		"gohan": 2044,
	}
	fmt.Println("lookup map:", lookup)		
}

func map_iterate() {
	lookup := map[string]int{
		"goku": 9001,
		"gohan": 2044,
	}
	// Iteration over maps isn’t ordered.
	// Each iteration over a lookup will return the key value pair in a random order.
	for key, value := range lookup {
		fmt.Println(key, value)
	}
}