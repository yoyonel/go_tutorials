package main

import "fmt"

var (
	name       = "Earth"
	desc       = "Planet"
	radius     = 6378
	mass       = 5.972E+24
	active     = true
	satellites = []string{
		"Moon",
	}
)

func main() {
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}
