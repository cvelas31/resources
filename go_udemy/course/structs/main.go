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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	fmt.Printf("%+v", jim)
	jim.updateName("Jimmy")
	jim.print()
	var alejo person
	alejo.firstName = "Alejo"
	alejo.lastName = "Anderson"
	fmt.Println(alejo)
	fmt.Printf("%+v", alejo)
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
