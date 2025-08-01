package main

import "fmt"

// func add(num1 int, num2 int) {
// 	sum := num1 + num2
// 	fmt.Println(sum)
// }

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumber(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func seyHello(name string) {
	fmt.Println("This is a", name)
}

func main() {
	// var a int
	// var b int = 10  // declaration with initialization
	// var c = "hello" // type inferred
	name := true
	name = false
	// name = "this"
	const pi = 3.14
	const greeting = "Hello"
	var x, y int = 1, 2
	var (
		a = 10
		b = "world"
		c = true
	)
	age := 30
	if age >= 50 {
		println("this is")
	} else {
		println("this is not")
	}

	fmt.Println(name)
	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	C := 30
	d := 40
	sum := add(C, d)
	println(sum)
	sum1, mul := getNumber(C, d)
	println(sum1)
	println(mul)
	seyHello("Rony")
}
