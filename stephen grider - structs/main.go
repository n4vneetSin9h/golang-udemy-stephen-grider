package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) printName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
	// p.firstName = firstName
}

func main() {
	john := person {
		firstName: "John", 
		lastName: "Doe",
		contactInfo: contactInfo {
			email: "john.doe@gmail.com",
			zipCode: 123456,
		},
	}
	fmt.Printf("%+v\n", john)

	john.printName()

	johnII := &john
	johnII.updateFirstName("John II")
	// john.updateFirstName("John II")

	john.printName()
}