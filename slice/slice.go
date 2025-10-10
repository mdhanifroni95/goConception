package main

import "fmt"

//Stack Allocation
// func stackExample() {
// 	s := []int{1, 2, 3}
// 	fmt.Println(s)
// }

//heap Allocation
func heapExample() []int {
	s := []int{10, 34, 21}
	fmt.Printf("Inside Function:\n")
	fmt.Printf("Address of slice variable s: %p\n", &s)
	fmt.Printf("Address of underlying array: %p\n", &s[0])
	return s
}

func main() {
	arr := [6]string{"this", "is", "a", "interview", "questions"}
	fmt.Println("arr", arr)

	// 1. Direct literal দিয়ে
	num := []int{10, 20, 30, 40}
	fmt.Println("num", num)

	// 2. make দিয়ে (len, capacity)
	s := make([]int, 3, 20)
	fmt.Println("s", s, "length", len(s), "cap", cap(s))
	s[0] = 100
	s[1] = 200
	s[2] = 300

	slice1 := arr[0:3]
	fmt.Println("slice1", slice1, "length", len(slice1), "cap", cap(slice1))

	slice1 = append(slice1, "read")
	slice1 = append(slice1, "write")
	slice1 = append(slice1, "code")
	slice1 = append(slice1, "debug")
	slice1 = append(slice1, "test")
	slice1 = append(slice1, "deploy")
	fmt.Println("slice1 app", slice1, "length app", len(slice1), "cap app", cap(slice1))
	fmt.Println("modify arr", arr, "length", len(arr), "cap", cap(arr))

	slice2 := slice1[4:6]
	fmt.Println("slice2", len(slice2), "capacity slice2", cap(slice2), "pinter", "slice1 length", len(slice1), "capacity slice1", cap(slice1))
	fmt.Printf("Address of slice variable Array: %p\n", &arr)
	fmt.Printf("Address of underlying array Array: %p\n", &arr[0])

	fmt.Printf("Address of slice variable slice1: %p\n", &slice1)
	fmt.Printf("Address of underlying array slice1: %p\n", &slice1[0])

	fmt.Printf("Address of slice variable slice2: %p\n", &slice2)
	fmt.Printf("Address of underlying array slice1: %p\n", &slice2[0])

	//nested slice literal

	/**
	matrix := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	fmt.Println("matrix", matrix)
	*/

	// Stack Allocation
	// stackExample()

	// Heap Allocation (Escape)
	h := heapExample()
	fmt.Printf("\nInside Main:\n")
	fmt.Printf("Address of slice variable h: %p\n", &h)
	fmt.Printf("Address of underlying array: %p\n", &h[0])

	// s[0] = 500
	// s = append(s, 400) //runtime error: index out of range [3] with length 3
	// fmt.Println("s", s, len(s), cap(s))
	// s1 := arr[1:4]
	// fmt.Println("s1", s1)

	// s2 := s1[1:2]
	// fmt.Println("s2", s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	//slice littler
	// slice := []int{1, 2, 3}
	// fmt.Println("slice", slice, "len:", len(slice), "cap:", cap(slice))

	//make function into initial slice
	// sliceMak := make([]int, 3, 5)
	// sliceMak[0] = 6
	// sliceMak[2] = 9
	// fmt.Println("sliceMak", sliceMak, len(sliceMak), cap(sliceMak))
}

/*
	2 Phases:
		1.compilation phase
		2.execution phase

		******* compile phase****

		** Code Segment **
		main = func() {...}
*/
