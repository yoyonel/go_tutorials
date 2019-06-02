# A First Step in Go

"Go is an attempt to combine the safety and performance of a statically-typed language
with the expressiveness and convenience of a dynamically-typed interpreted language."

Ideas behing Go:
...
- Statically typed for compilation and runtime safety
...
- A garbage collector for safe and automatic memory management

Notion de 'workspace' a bien comprendre.
Ca semble être une sorte d'organisation (layout) en terme de structures de fichiers (directories).
A préciser.

## Composite types
Composite types: array, slice, map, ...

named type: sorte de typedef struc C
```go
type amu float64

type metalloid struct {
    name string
    number int32
    weight amu
}
```

## Methods, Interfaces, and Objects.
"Go types do not use a class hierarchy to model the world as is the wase with other object-oriented languages."

## Concurrency and channels

goroutine: `go` keyword
and `<-` for sending result into channels. Notion of channel `make(chan <type>, <buffer_size>)` 

goroutine run independently and uses the Go channels to communicate and coordinate (the calculation and the final result).

=> intéressant comme pattern ^^

## Memory management and safety

"Pointer arithmetic is not permitted at runtime; ..."

# Testing and code coverage

"Built-in API and tools designed specially for automated testing, benchmarking, and code coverage."
"test tools use simple conventions to automatically inspect and instrument the test functions found in your code."
"In a separate sourec file, [...]"
`go test .`
```
Running tool: /usr/bin/go test -timeout 30s -run ^(TestDivide)$

PASS
ok  	_/home/latty/Prog/__GOLANG__/go_tutorials/Learning_Go_Programming/ch01/testexample	0.070s
Success: Tests passed.
```
