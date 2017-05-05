package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

func main() {
	declarations_initializations_1()
	declarations_initializations_2()
	declarations_initializations_3()
}

func declarations_initializations_1() {
	goku := Saiyan{
		Name: "Goku",
		Power: 9000,
	}
	fmt.Println(goku)
}

func declarations_initializations_2() {
	goku := Saiyan{}
	fmt.Println(goku)
}

func declarations_initializations_3() {
	goku := Saiyan{ Name: "Goku"}
	goku.Power = 9001
	fmt.Println(goku)
}

func declarations_initializations_4() {
	goku := Saiyan{"Goku", 9002}
	fmt.Println(goku)
}