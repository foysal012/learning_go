/// First Program

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

/// Data Type declearation

// package main

// import "fmt"

// func main() {

// 	// variable type 1
// 	var a int = 10
// 	fmt.Println(a)

// 	// variable type 2
// 	var b float32 = 10.32
// 	fmt.Println(b)

// 	// variable type 3
// 	var c bool = false
// 	fmt.Println(c)

// 	// variable type 4
// 	var d string = "Foysal"
// 	fmt.Println(d)

// 	// variable type 5
// 	e := 105
// 	fmt.Println(e)

// 	// variable type 6
// 	f := 105
// 	f = 150
// 	fmt.Println(f)

// 	fmt.Print(f, e, d, c, b, a)

// }

/// Switch & If...Else

// <, <=, >, >=, ==
// &&, ||, !

// package main

// import "fmt"

// func main() {
// 	// age := 18.5

// 	// if age > 18 {
// 	// 	fmt.Println("You are married")
// 	// } else if age < 18 {
// 	// 	fmt.Println("You are a Teenager")
// 	// } else if age == 18 {
// 	// 	fmt.Println("You are a candidate for marriage")
// 	// } else {
// 	// 	fmt.Println("You are a child")
// 	// }

// 	// age := 16
// 	// gender := "female"
// 	// isSweety := false

// 	// if age == 18 && gender == "male" {
// 	// 	fmt.Println("Your Current age ready for married")
// 	// } else if age >= 18 || gender == "male" {
// 	// 	fmt.Println("You are already married")
// 	// } else if !isSweety {
// 	// 	fmt.Println("Married is a sweet")
// 	// }

// 	a := 3

// 	switch a {
// 	case 1:
// 		fmt.Println("Your Entering number is 1")
// 	case 2, 3:
// 		fmt.Println("The Number is Either 2 or 3")
// 	default:
// 		fmt.Println("Ther Number is neither 1 nor 2 or 3")
// 	}

// }

/// function

// package main

// import "fmt"

// func add(num1 int, num2 int) {
// 	sum := num1 + num2
// 	fmt.Println("Result is: ", sum)
// }

// func main() {

// 	a := 10
// 	b := 20

// 	// sum := a + b
// 	// fmt.Println("Result is: ", sum)

// 	add(a, b)
// 	fmt.Println("Math is remining")

// 	add(5, 7)
// 	fmt.Println("Math is complete")

// }

/// function with return Value

// package main

// import "fmt"

// func add(num1 int, num2 int) int {
// 	sum := num1 + num2

// 	return sum
// }

// func main() {
// 	a := 10
// 	b := 20

// 	sum := add(a, b)

// 	fmt.Println(sum)
// }

// func add(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2

// 	mul := num1 * num2

// 	return sum, mul
// }

// func main() {
// 	a := 10
// 	b := 20

// 	p, q := add(a, b)

// 	fmt.Println(p)
// 	fmt.Println(q)
// }

/// More Function

// package main

// import "fmt"

// func PrintName() {
// 	fmt.Println("Education is must!")
// }

// func PrintName2(name string) {
// 	fmt.Println("You are the mr. ", name)
// }

// func main() {
// 	PrintName()
// 	PrintName2("Foysal Joarder")

// }

/*Day-2*/

/// All Functions

// package main

// import "fmt"

// func welcomeMessage() {
// 	fmt.Println("Welcome to the application")
// }

// func getname() string {
// 	var name string
// 	fmt.Println("Please Enter your Name:")
// 	fmt.Scanln(&name)
// 	return name
// }

// func getTwoNumber() (int, int) {
// 	var num1 int
// 	var num2 int

// 	fmt.Println("Please Enter your First Number:")
// 	fmt.Scanln(&num1)

// 	fmt.Println("Please Enter your Last Number:")
// 	fmt.Scanln(&num2)

// 	return num1, num2
// }

// func addNumber(num1 int, num2 int) int {
// 	sum := num1 + num2

// 	return sum
// }

// func displayScreen(name string, sum int) {
// 	fmt.Println("Hello, Mr. ", name)
// 	fmt.Println("Summation is ", sum)
// }

// func goodByeMessage() {
// 	fmt.Println("Thank you for using our application")
// 	fmt.Println("Good bye")
// }

// func main() {

// 	// fmt.Println("Welcome to the application")

// 	// var name string
// 	// fmt.Println("Please Enter your Name:")
// 	// fmt.Scanln(&name)

// 	// var num1 int
// 	// var num2 int
// 	// fmt.Println("Please Enter your First Number:")
// 	// fmt.Scanln(&num1)

// 	//fmt.Println("Please Enter your Last Number:")
// 	//fmt.Scanln(&num2)

// 	// sum := num1 + num2

// 	// fmt.Println("Hello, Mr. ", name)
// 	// fmt.Println("Summation is ", sum)

// 	// fmt.Println("Thank you for using our application")
// 	// fmt.Println("Good bye")

// 	// calling all function serilally
// 	welcomeMessage()
// 	name := getname()
// 	num1, num2 := getTwoNumber()
// 	sumation := addNumber(num1, num2)
// 	displayScreen(name, sumation)
// 	goodByeMessage()
// }

/// What is Scope

package main

import "fmt"

// Global portion
var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println("Summation:", z)
}

func main() {

	p := 40
	q := 50

	add(p, q)

	add(a, b)

	add(p, b)

	add(q, a)
}

/*Day-3*/

/// scope more
