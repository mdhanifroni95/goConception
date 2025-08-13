package main

import "fmt"

// var (
// 	a = 10
// 	b = 40
// )

// func add(num1 int, num2 int) {
// 	res := num1 + num2
// 	printNum(res)
// }

func main() {
	func(a int, b int) {
		c := a + b
		fmt.Println("This is an anonymous function in scop package", c)
		printNum(c)
	}(9, 8)
}

func printNum(res int) {
	fmt.Println("This is a function in scop package", res)
}
