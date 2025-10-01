package main

import (
	"fmt"
	"math"
)

/*
	- If-else
	- Switch case
	- Function
*/

func runIfElse() {
	// if-else
	num := 34

	if num == 34 {
		fmt.Println("Yeah boy")
	} else if num != 34 && (num > 30 || num <= 39) {
		fmt.Println("Holy")
	} else {
		fmt.Println("please")
	}

	// short if
	lim := 225.0
	// v is in if only
	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}

	// switch case
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Printf("Today is %v", day)
	}

	// switch with logic in case (no variable in switch)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}

func add(x float64, y float64) (float64, string) {
	fmt.Println("Add this", x, y)
	v := x + y
	return v, "good boy"
}

// naked return
func split(sum int) (x int, y int) {
	x = sum * 4
	y = sum - x
	return
}

// signature of function

// higher order function

// กูหลุด เดี๋ยวกลับมาเขียน
