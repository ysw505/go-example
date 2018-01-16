package main

import (
	"fmt"
	"errors"
)

func echo(keyword string) ( string, error ) {
	if keyword == "error" {
		return "" , errors.New("Error !!")
	} else if keyword == "panic" {
		panic("panic") // panic make abort
	} else {
		return keyword, nil
	}
}

func main() {
	_ , err := echo("panic")

	if err != nil {
		fmt.Println(err.Error())
	}
}