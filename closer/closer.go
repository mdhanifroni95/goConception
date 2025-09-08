package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Print(money)
	}
	return show
}

func main() {
	call()
}

func call() {
	incr1 := outer()

	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func init() {
	fmt.Println("=== Bank ===")
}

/*
	Phases:
		1. compilation phase (compile time)
		2. execution phase (run time)

		********* Compilation phase **********
		 *** code Segment ***
		 a = 10
		 outer= func(){}
		 outerAnonyms1=fun(){}
		 call= func(){}
		 main= func(){}
		 init= func(){}

		 go run main.go => compile it =>./main
		 go build main.go => compile it =>main
		 ./

*/

/**
  *** Binary Executable File ***
	main
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
	0010111011101110001
*/
