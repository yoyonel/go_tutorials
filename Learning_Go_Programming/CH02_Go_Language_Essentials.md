# Chapter 02: Go Language Essentials

## The Go source file

3 main sections in a typical Go source file:
- The Package Clause: `package main`
- The Import Declaration: `import "fmt"`
- The Source Body: `func main() { ... }`

"It's consider an error (compilation) to include an unused package [...]."

"All Go source files must end with the .go suffix."

Simple idioms and conventions:
- Optional semicolon: at the end of a sentence `;`
- Multiple lines: ~ close to Python rules

tools: `gofmt` `govet`: analyzing code for structural problems with code elements.

## Go identifiers

"The first position of an identifier must be a letter or an underscore"
"camel case"

"basic rule: 'you declare it, you must use it.'"
Blank identifier `_` (like Python)

### muting package
`_` can be used for muting package import like: `import _ "log"`
(notion de "package initialization lifecycle")

### muting unwanted function results
`_, execFile := filepath.Split("/opt/data/bigdata.txt")`
filepath.Split return two values: first is the parent path (`/opt/data`), and the second is the file name (`bigdata.txt`).
Other uses of this idiom can be error-handling and `for` loops.

### Built-in identifiers

- [Type] Numeric: `byte, int, int8, int16, rune, complex64, uintptr, ...`
- [Type] Error: `error`
- [Values] Constant counter: `iota`
- [Values] Uninitialized value: `nil`
- [Functions] Initialization: `make(), new()`
- [Functions] Complex numbers: `complex(), imag(), real()`
- [Functions] Error Handling: `panic(), recover()`

### Go variables

- Variable declaration: `var <identifier list> <type>`
- Zero-value for types: Interface, function, channel, slice, map and pointer => `nil`
- Initialized declaration: `var <identifier list> <type> = <value list or initializer expression>`
  examples: `var salletites = []string{"Moon", }`, `var name, desc string = "Earth", "Planet"`
- Omitting variable types: `var <identifier list> = <value list or initializer expression>`
  -> Notion of 'Inferred type':
  - Complex Numbers: (`-5.0i`,`3i`, `(0+4i)`)  inferred to `complex128`
  - Array values: `[2]int{-76, 8080}` inferred to `array` (`[2]int`)
  - Map values: `map[string]int{"Sun": 685800, "Earth": 6378}` inferred to `map[string]int`
  - Slice values: `[]int{-76, 0, 1244, 1840}` inferred to `slice` (`[]int`)
- Short variable declaration: `name := "Neptune"` (no more `var` or variable type)
    - can only be used within a function block
    - `:=` cannot be used to update a previously declared variable (update using `=` sign)

#### Variable scope and visibility

"Go uses lexical scoping based on code blocks to determine the visibility of variables within a package. [...]
As a general rule, a variable is only accessible from within the block where it is declared and visible to all nested sub-blocks."

"When a variable is declared at package level (outside of a function or method block), it is globally visible to the entire package,
not just to the source file where the variable is declared. [...]"

#### Variable declaration block

```golang
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
```

### Go Constant

`const <identifier list> type = <value or initializer expressions>`
examples:
```golang
const a1, a2 string = "Mastering", "Go"
[...]
const h time.Duration = 4 * time.Second
```
(with types)

### Untyped constants

`const <identifier list> = <value or initializer expressions>`
examples:
```golang
const i = "G is" + " for GO "
[...]
const o = time.Millisecond * 5
```

### Assigning untyped constants

```golang
const m2 = 1.414213562373095048801688724209698078569671875376...
var u1 float32 = m2
var u2 float64 = m2
u3 := m2
```
m2 stored raw untyped value
=> type of u3 is `float64`

### Constant declaration block
Same idea from 'variable declaration block':
```go
const (
    a1, a2 string = "Mastering", "Go"
    [...]
    h time.Duration = 4 * time.Second
)
```

### Constant enumeration

"One interesting usage of constants is to create enumerated values. [...]
assign the pre-declared constant value `iota` to a constant identifier in the declaration block, [...]":
```go
const (
    StartHyperGiant = iota
    [...]
    StarBrownDwarf
)
```
=> `iota` assign to 0

- Overriding the default enumeration type: 
```go
const (
    StarDwarf byte  = iota
    [...]
    StarBrownDwarf
)
```

- Using `iota` in expressions
"When `iota` appears in an expression, the same mechanism work as expected. The compiler will apply the expression
for each sucessive increaging value of `iota`.":
```go
const (
    StarHyperGiant = 2.0*iota
    StarSuperGiant
    [...]
    StarSubGiant
)
```
results:
```go
    StarHyperGiant = 0  [float64]
    StarSuperGiant = 2  [float64]
    [...]
    StarSunGiant = 8    [float64]
```

- Skipping enumerated values
```go
const(
    _               = iota      // value 0
    StarHyperGiant  = 1 << iota // value 2
    [...]
    _                           // value 64
    StarDwarf                   // value 128
    [...]
)
```

### Go Operators

`++` `--` increment/decrement operators:
```go
func reverse(s string) {
    for i := len(s) - 1; i >= 0; {
        fmt.Print(string(s[i]))
        i --
    }
}
```
"increment and decrement operators are statements, not expressions":
```go
nextChar := i++     // syntax error
fmt.Println("Current char", i--)    // syntax error
nextChar++          // Ok
```

- Bitwise operators: 
    - `&^`: Bitwise AND NOT
    - `^a`: Unary bitwise complement
