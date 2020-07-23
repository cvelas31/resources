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
- Struct syntax:




# Execute
- Run:
go run main.go

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
func name(arguments){
    Body
}
```

**Structs**
