# Go

# Summary
- Is a pass by value language (Every input that enters to function copies to another memory slot)
- At every file on our working dir must have `package name``
- Imports at the beginning as strings
- Custom types created as `type name_type []string`
- Function syntax:
    ```go
    func (type_owner or receiver) name_func(input type_input) (output_type) {
        body
        }
    ```
- Struct syntax `type name struc (key: value)`
- Operand `&` take the adress of that variable (Point the address)
- Operand `*`give the value a memory adress is pointing
- Map (key, value). Keys one type and values another type.
- Interfaces. All types inside that have some "same" functions (return and  get inputs of the same type) and its called the same.
    - Its kind of shared type, inherits the functions of it.


# Execute
- Run:
go run main.go other.go

## CLI
1. go run
    - Compiles and executes one or N files
2. go build
    - Compiles a bunch of go source code files
    - Then we can run the name of the compiled files as .exe in Windows
3. go fmt
4. go install
5. go get
6. go test
    - To test some code

## Definitions
**Package**: package means project or  workspace, can have many files .go on it
The first line must declare `package main`
There are executable packages and reusable packages (Libraries)

`package main`: create a executable file (Needs main function)
`package blabla`: Creates reusable code

**Imports**
Access to other packages. (fmt for examples is terminal or format things)
Standard packages:
- math
- format
- crypto
- ioddd
- debug
- encoding

**Functions**
Regular functions
```Go
func name(arguments) outputType{
    Body
    return outputs
}
```

**Pointers**
```go
func (pointerToPerson *person) updateName(){//This pointer is a type description: Means we are working with a ponter to a person 
    *pointerToPerson//This is an operator: Means we want to manipulate the value the pointer is referencing
}
```

**Maps**
```go
a = map[string]string{
    "Hello": "Hola",
    }
a["Como"] =  "How"
```

**Interfaces**
```go
type ExampleInterface interface {
    exampleFunction(string) string
    exampleFunction2(*pointer) int
}
```

**Type Assertion**
```go
var i interface{} = "hello"
s := i.(string)
fmt.Println(s)
```

## Types in Go
- **Value Types**: int, float, string, bool, structs. This are kind of nonmutable which copies when entering to a function. Need to worry about pointers *.

- **Reference Types**: slices, maps, channels, pointers, functions. This are kind of lists, dictiionaries which gets the same object or pointer in memory when copied.