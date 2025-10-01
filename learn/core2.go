package main

import (
	"fmt"
)

func runArray() {
	// var array [lenght]type{val}
	// nums := [...]int{1, 2, 3, 4, 5}

	test := [...]string{"go", "is", "fun"}

	// index , value
	for i, v := range test {
		fmt.Println(i, v)
	}
}

func runLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}
	// while loop
	sum := 1
	for sum < 5 {
		sum += sum
		fmt.Println("sum: ", sum)
	}
}
