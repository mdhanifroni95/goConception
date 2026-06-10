package main

import "fmt"

func sayHi(name string) {
	fmt.Println(name)
}

// Return value সহ function

// func add(a int, b int) int {
// 	return a + b
// }

// short syntax
// func add(a, b int) int {
// 	return a + b
// }

func square(n int) int {
	return n * n
}

func main() {
	fmt.Println(square(19))
	fmt.Println(square(40))
	fmt.Println(square(10))
}
