package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Your are Adult")
	} else if age <= 18 && age >= 13 {
		fmt.Println("Your are Teenager")
	} else {
		fmt.Println("Your are Child")
	}
}
