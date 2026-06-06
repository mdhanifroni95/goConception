package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Join", Age: 9}
	fmt.Println("Name", p1.Name, "Age", p1.Age)
	var p2 Person
	p2.Name = "Tom"
	p2.Age = 10
	fmt.Println(p2)

	p3 := &Person{Name: "Hassan", Age: 40}
	fmt.Println(p3.Name)
	p3.Age = 45
	fmt.Println(p3.Age)
}
