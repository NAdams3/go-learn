package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type Manager struct {
	Employee
	Test Employee
}

func (e Employee) printFullName() {
	fmt.Printf("%s %s\n", e.FirstName, e.LastName)
}

func employee() {
	m := Manager{
		Test: Employee{
			FirstName: "Jane",
			LastName:  "Doe",
		},
		Employee: Employee{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	n := Manager{
		Test: Employee{
			FirstName: "Jim",
			LastName:  "Picard",
		},
	}

	fmt.Printf("manager: %+v \n", m)

	m.printFullName()
	m.Employee.printFullName()
	m.Test.printFullName()
	fmt.Printf("m firstname: %v\n", m.FirstName)

	n.printFullName()
	n.Employee.printFullName()
	n.Test.printFullName()
}
