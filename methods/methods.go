package methods

import (
	"fmt"
	"time"
)

type Named interface {
	FirstName() string
}

type EmployeeId int

type Employee struct {
	ID        EmployeeId
	firstName string
	LastName  string
	DateHired time.Time
}

func NewEmptyEmployee() *Employee {
	return &Employee{
		ID: -1,
	}
}

// Examples of getters and setter fields

func (e *Employee) FirstName() string {
	return e.firstName
}

func (e *Employee) SetFirstName(value string) {
	e.firstName = value
}

type Manager struct {
	Employee
	Reports []EmployeeId
}

// overloaded method
func (m *Manager) FirstName() string {
	return "Manager " + m.Employee.FirstName()
}

func example1() {
	e1 := &Employee{
		ID:        1,
		firstName: "John",
		LastName:  "Doe",
		DateHired: time.Now(),
	}

	fmt.Println("e1 first name", e1.FirstName())
	fmt.Println("e1 date hired", e1.DateHired)

	fmt.Println("e1 first name", e1.FirstName())
	e1.SetFirstName("Dan")
	fmt.Println("e1 first name", e1.FirstName())

	emptyEmployee := NewEmptyEmployee()
	fmt.Println("empty employee", emptyEmployee.ID)

	mgr1 := &Manager{
		Employee: Employee{ID: 1, firstName: "John"},
	}

	doSomethingWithEmployee(e1)
	// Does not work :(
	// doSomethingWithEmployee(mgr1)

	// but this does?
	doSomethingWithEmployee(&mgr1.Employee)
	fmt.Println(mgr1.FirstName())

	// also this is interesting
	fmt.Println("== Named interface")
	doSomethingWithNamed(e1)
	doSomethingWithNamed(mgr1)
	doSomethingWithNamed(&Employee{
		ID:        99,
		firstName: "Anonymous",
	})

}

func doSomethingWithEmployee(e1 *Employee) {
	fmt.Println(e1.FirstName())
}

func doSomethingWithNamed(n1 Named) {
	fmt.Println(n1.FirstName())
}

func RunExamples() {
	example1()
}

func RunChallanges() {
}
