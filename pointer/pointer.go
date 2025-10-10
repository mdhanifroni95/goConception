package main

import "fmt"

type Student struct {
	Name string
	Age  int
	dob  string
}

func print(arr *Student) {
	fmt.Println("arr:", arr)
}
func changeValue(num *int) {
	*num += 10
	fmt.Println("num:", num)
}
func main() {
	x := 42
	p := &x
	*p = 21
	// arr := []int{1, 2, 3}
	changeValue(&x)
	fmt.Println("x:", x)
	std := Student{Name: "Rimsha", Age: 6, dob: "01-01-2018"}
	print(&std)
	fmt.Println(x)
}
