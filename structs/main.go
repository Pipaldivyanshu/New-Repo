package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//name := person{firstName: "Alex", lastName: "Baker"}

	//fmt.Println(name)

	//var name person

	//name.firstName = "ALex"
	//name.lastName = "Baker"

	//fmt.Println(name)

	jim := person{

		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jimHalpert@gmail.com",
			zipCode: 1237,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newName string) {

	(*pointerToPerson).firstName = newName
}

func (p person) print() {

	fmt.Printf("%+v", p)

}
