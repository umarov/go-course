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
	muzafar := person{
		firstName: "Muzafar",
		lastName:  "Umarov",
		contactInfo: contactInfo{
			email:   "email@email.com",
			zipCode: 90909,
		},
	}

	muzafar.print()

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Coder"
	alex.contactInfo.email = "test@email.com"
	alex.contactInfo.zipCode = 44232

	alex.print()

	alex.updateName("Alecks")
	alex.print()
}

func (person *person) updateName(newFirstName string) {
	person.firstName = newFirstName
}

func (person person) print() {
	fmt.Printf("%+v \n", person)
}
