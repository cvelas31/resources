package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//Form 1
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.print()
	jim.updateName("Jimmy")
	jim.print()
	//Form 2
	var alejo person
	alejo.firstName = "Alejo"
	alejo.lastName = "Anderson"
	fmt.Println(alejo)
	fmt.Printf("%+v", alejo)

	//Form 3
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//Form 4 with slicing and pointers
	name := "Bill"
	nameP := &name
	fmt.Println(nameP)
	print(name)
}

func print(name string) {
	fmt.Println(&name)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
