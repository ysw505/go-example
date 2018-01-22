package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	arrInt := []int{}
 	rand.Seed(time.Now().UnixNano())

	for i := 0 ; i < 100 ; i++ {
		arrInt = append(arrInt, rand.Int() % 100)
	}

	newArr := arrMap(arrInt, func(each int) bool {
		if each < 10 {
			return true
		}

		return false
	})

	fmt.Println(newArr)
}

func arrMap(sourceArr []int, eachFunc func(each int) bool) []int {

	newArr := []int{}

	for _, v := range sourceArr {
		if eachFunc(v) {
			newArr = append( newArr , v )
		}
	}

	return newArr
}