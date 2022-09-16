package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	ryan := Employee{
		name:   "Ryan Morrrison",
		salary: 50000,
		contactInfo: Contact{
			email:   "morrison@rnt.com",
			address: "West Nusa Tenggara, Indonesia",
			phone:   001122334455,
		},
	}

	fmt.Printf("%+v\n", ryan)
	fmt.Printf("Employees's email: %s\n", ryan.contactInfo.email)

	ryan.contactInfo.email = "morrisey-new@rnt.com"
	fmt.Printf("Employee's new email: %s\n", ryan.contactInfo.email)

	myContact := Contact{
		email:   "rianto@rnt.com",
		phone:   464316461318005,
		address: "Sumbawa",
	}

	fmt.Println(myContact)
}
