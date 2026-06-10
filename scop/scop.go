package main

// Scope in Go refers to the visibility and lifetime of variables, functions, and other identifiers. It determines where these identifiers can be accessed and modified within a program. Go has several types of scopes, including package scope, function scope, block scope, and file scope.
import "fmt"

// var (
// 	a = 10
// 	b = 40
// )

// func add(num1 int, num2 int) {
// 	res := num1 + num2
// 	printNum(res)
// }

// func main() {
// 	func(a int, b int) {
// 		c := a + b
// 		fmt.Println("This is an anonymous function in scop package", c)
// 		printNum(c)
// 	}(9, 8)
// }

// func printNum(res int) {
// 	fmt.Println("This is a function in scop package", res)
// }

//package scope
var appName = "My App"

func main() {
	//function scope
	name := "John Doe"

	//block scope
	if age := 20; age > 18 {
		fmt.Println(age)
	}
	// fmt.Println(age) // This will cause an error because age is not accessible outside the if block

	fmt.Println(name)
	fmt.Println(appName)
}
