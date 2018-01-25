package main

import "fmt"


type Company interface {
	PrintCompanyInfo()

}


type Naver struct {
	Name string
	EmployeeNo string
}

func (n Naver) PrintCompanyInfo() {
	fmt.Println(n.EmployeeNo, n.Name)
}

type Nbp struct {
	Name string
	EmployeeNo string
}


func (n Nbp) PrintCompanyInfo() {
	fmt.Println(n.EmployeeNo, n.Name)
}


func PrintCompany(c Company) {
	c.PrintCompanyInfo()

}

func main() {
	PrintCompany(Naver{"123","123"})
	PrintCompany(Nbp{"123","123"})
}


