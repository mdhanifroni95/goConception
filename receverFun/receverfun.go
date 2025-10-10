package main

import "fmt"

// Example of a method with a value receiver
// type Person struct {
// 	name string
// 	age  int
// }

// func (per Person) display() {
// 	fmt.Println("Hello", per.name)
// }

// func main() {
// 	p1 := Person{name: "Rimsha", age: 06}
// 	p1.display()
// }

// Example of a method with a pointer receiver
// type Counter struct {
// 	value int
// }

// func (c *Counter) increment() {
// 	c.value++
// }

// func main() {
// 	c := Counter{value: 10}
// 	c.increment()
// 	fmt.Println("Counter value:", c.value)
// }
// Example of a Mixed Receiver

type Spae struct {
	width, hight int
}

//value receiver (read-only)
func (s Spae) area() float64 {
	return float64(s.width) * float64(s.hight)
}

func (s *Spae) scale(factor float64) {
	s.width *= int(factor)
	s.hight *= int(factor)
}

func main() {
	rect := Spae{width: 10, hight: 5}
	fmt.Println("Area:", rect.area())

	rect.scale(2)
	fmt.Println("Scaled Area:", rect.area())
}
