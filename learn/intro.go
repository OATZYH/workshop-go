package main

import "fmt" // standard I/O package like stdio in c

/*
	- Declare variable
	- Formatting
	- Zero value
*/
// Declare outside
var out = "Outside"

func runIntro() { // Go auto insert semi-colon can't drop this bracket
	// run the program with "go run hello.go"
	// no semi-colon at the end of the line
	fmt.Println("Hello, World!")
	// print("Hello")    // ไม่มี newline
	// println("Hello")  // มี newline ต่อท้าย

	// declare variable
	var x string = "Declare"
	fmt.Println(x)
	fmt.Println(out)

	// Not allow to run var that not use

	msg := "test" + "2"
	fmt.Println("message", msg)

	// var and :=
	// comment = "comment" this error because not initializes
	// SO we can do 2 things
	// var comment = "string" OR
	// comment := "string" > short version auto type

	/*
		Output:
		topic: Avengers: Endgame
		year: 2019
		rating: 8.4
		category: Sci-Fi
		superhero: true

		1. Print เพียวๆ
		2. ใช้ ตัวแปร
		3. Raw string
	*/

	// var t1, t2, t22, t3 = "string", 1, 1.5, true
	// fmt. Printf("t1: %s \n", t1)
	// fmt. Printf("t2: %d\n", t2)
	// fmt.Printf("t22: %.2f\n", t22)
	// fmt.Printf("t3: %t\n", t33)

	// %#v don't know type
	// %T find type

	r := rune('🙏')        // infer type จาก cast ว่าเป็น rune
	fmt.Println(r)        // แสดงค่า int32 ของ 🙏
	fmt.Printf("%c\n", r) // แสดงเป็นตัวอักษร 🙏

	/*
		Output:
		topic: Avengers: Endgame
		year: 2019
		rating: 8.4
		category: Sci-Fi
		superhero: true
		favorite: ⭐
	*/

	// Zero value
	// var msg string
	// var age int
	// var price float64
	// var check bool
	// var r rune

	// fmt.Printf("msg: %s\n", msg)
	// fmt.Printf("age: %d\n", age)
	// fmt.Printf("price: %.2f\n", price)
	// fmt.Printf("check: %t\n", check)
	// fmt.Printf("r: %c\n", r)
}
