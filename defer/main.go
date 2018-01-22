package main

import "fmt"

func main() {
	fmt.Println("Defer is stack")

	defer func() {
		fmt.Println("defer 0")
	}()

	foo()
}


func foo() {
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		fmt.Println("defer 2")
	}()

	defer func() {
		fmt.Println("defer 3")
	}()

	defer func() {
		fmt.Println("defer 4")
	}()
}