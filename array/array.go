package main

import "fmt"

var arr []int
var stringArr []string

func main() {
	// stringArr := []string{}
	num := [2]int{34, 78}
	fmt.Println(num)
	arr = append(arr, 45)
	fmt.Println(arr)
	arr = append(arr, 98, 56, 23)
	fmt.Println(arr)
	stringArr = append(stringArr, "hello", "world")
}
