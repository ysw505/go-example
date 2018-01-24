package main

import "fmt"

type Employee struct {
	Name string
	EmpNo string
}

func (e Employee) Info() {
	fmt.Println(e.EmpNo, e.Name)
}

func NewEmployee(empNo, name string) Employee {
	return Employee{name, empNo}
}

func NewEmployeePointer(empNo, name string) *Employee {
	return &Employee{name, empNo}
}

func PrintEmployeeInfo(e Employee) {
	e.Info()
}

func ChangeEmployeeInfo(e Employee) {
	e.Name = "mutable"
}

func ChangeEmployeeInfoPtr(e *Employee) {
	e.Name = "mutable"
}

func main() {

	a := NewEmployee("NB11111", "YSW")

	b := NewEmployeePointer("NB12345", "DGW")

	fmt.Println("before change")
	PrintEmployeeInfo(a)
	PrintEmployeeInfo(*b)

	fmt.Println("after change")
	ChangeEmployeeInfo(a)
	ChangeEmployeeInfoPtr(b)
	PrintEmployeeInfo(a)
	PrintEmployeeInfo(*b)

}
