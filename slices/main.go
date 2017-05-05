package main

import(
	"fmt"
	"strings"
	"math/rand"
	"sort"
)

type Mapable interface {
}

// This function maps the every element for
// all []types of array.
func Map(mapper func(Mapable), list ... Mapable) {
    for _, value := range list {
        mapper(value)
    }
}

func main() {
	slice_1()
	slice_2()
	slice_3()
	slice_4()
	algo_slice_1()
	slice_5()
	slice_6()
	algo_slice_2()
	algo_slice_3()
	algo_slice_4()
	algo_slice_5()
}

func printTitle(name string) {
	// url: http://stackoverflow.com/questions/33139020/can-golang-multiply-strings-like-python-can
	name += "()"
	fmt.Printf("%s\n%s\n", name, strings.Repeat("-", len(name)))
}

func slice_1() {
	printTitle("slice_1")
	scores := []int{1, 4, 43, 9001}
	fmt.Println("scores: ", scores)
}

func slice_2() {
	printTitle("slice_2")
	scores := make([]int, 10)
	printSlice(scores)

	if len(scores) > 0 {
		scores[7] = 9033
		printSlice(scores)
	}		
}

func slice_3() {
	printTitle("slice_3")
	scores := make([]int, 0, 10)
	printSlice(scores)

	if len(scores) > 0 {
		scores[5] = 42
		printSlice(scores)
	}
}

func slice_4() {
	printTitle("slice_4")
	scores := make([]int, 0, 10)
	printSlice(scores)

	// re-slice our slice
	scores = scores[0:8]
	printSlice(scores)
	
	// re-slice our slice (same operation but without a assignation/declaration variable)
	fmt.Println("len(make([]int, 0, 10)[0:8]): ", len(make([]int, 0, 10)[0:8]))

	// panic: runtime error: slice bounds out of range
	// scores = scores[0:12]
}

func algo_slice_1() {
	printTitle("algo_slice_1")
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)
	printSlice(scores)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		printSlice(scores)
		//
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println("New cap=", c)
		}
	}
	// results:
	// 5
	// 10
	// 20
	// 40
}

// url: https://tour.golang.org/moretypes/11
func slice_5() {
	printTitle("slice_5")
	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. 

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice_String(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice_Bool(s []bool) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice_6() {
	printTitle("slice_6")

	// four common ways to initialize a slice
	names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10)
	var names2 []string
	scores := make([]int, 0, 20)

	printSlice_String(names)
	printSlice_Bool(checks)
	printSlice_String(names2)
	printSlice(scores)
}

func algo_slice_2() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores)
	// result expected: [1 2 999 4 5]
}

func algo_slice_3() {
	printTitle("algo_slice_3")
	haystack := "the spice must flow"
	fmt.Println("string 'haystack':", haystack)
	fmt.Println("Index of the first space in a string after the first five characters:", strings.Index(haystack[5:], " "))
}

func algo_slice_4() {
	scores := []int{1, 2, 3, 4, 5}
	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)
}

// start of an effecient way to remove a value from an unsorted slice
func removeAtIndex(sources []int, index int) []int {
	lastIndex := len(sources) - 1
	// swap the last value and the value we want to remove
	sources[index], sources[lastIndex] = sources[lastIndex], sources[index]
	return sources[:lastIndex]
}

// built-in function: copy (with slices)
func algo_slice_5() {
	printTitle("algo_slice_5")

	// declare the array
	scores := make([]int, 100)
	// put random values into the list/array
	for i := 0; i < 100; i ++ {
		scores[i] = int(rand.Int31n(1000))
	}
	// sort of the list/array
	sort.Ints(scores)

	// find the 5th lower values
	worst := make([]int, 5)
	// => get the 5th first values in the sorted array
	copy(worst, scores[:5])

	fmt.Println("5th lower values:", worst)

}