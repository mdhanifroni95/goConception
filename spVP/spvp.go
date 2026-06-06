package main

import "fmt"

func add(x int, y int) int {
	res := x + y
	return res
}

func main() {
	var a int = 10
	sum := add(a, 4)
	fmt.Println("sum is", sum)
}
