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

package main

import "fmt"

func PrintName() {
	fmt.Println("Education is must!")
}

func PrintName2(name string) {
	fmt.Println("You are the mr. ", name)
}

func main() {
	PrintName()
	PrintName2("Foysal Joarder")

}
