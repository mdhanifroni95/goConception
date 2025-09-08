package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{Name: "Alice", Age: 30}
	u2 := User{Name: "Rimsha", Age: 6}
	// fmt.Println("Name", u1.Name)
	// fmt.Println("Age", u1.Age)
	fmt.Println("u1", u1)
	fmt.Println("Name", u2.Name)
	fmt.Println("Age", u2.Age)
	// u2 := User{Name: "Rimsha", Age: 06}

	emp := struct {
		Id   int
		Role string
	}{Id: 1, Role: "Developer"}
	fmt.Println("emp", emp)
	u1.Greet()
	u2.Greet()
}

func (u User) Greet() {
	fmt.Println("Hello, my name is", u.Name)
}
