package main

import "fmt"

func main() {
	// day := "Friday"

	// basic switch
	// switch day {
	// case 1::
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// multiple values in one case

	// switch day {
	// case "Friday", "Saturday":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Working day")
	// }

	// switch work if else
	// marks := 75

	// switch {
	// case marks >= 80:
	// 	fmt.Println("Grade A")
	// case marks >= 60:
	// 	fmt.Println("Grade B")
	// case marks >= 40:
	// 	fmt.Println("Grade C")
	// default:
	// 	fmt.Println("Grade F")

	// }

	// Normally Go এক case থেকে পরের case-এ যায় না। কিন্তু চাইলে fallthrough ব্যবহার করা যায়।

	// num := 1
	// switch num {
	// case 1:
	// 	fmt.Println("One")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("Two")
	// default:
	// 	fmt.Println("Other number")
	// }

	// switch-এর ভিতরে variable declare

	switch day := 2; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
	}

}
