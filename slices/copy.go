package main

import "fmt"

func main() {

	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", array)
	s1 := array[8:]
	s2 := array[1:5]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", array)

	copy(s1, []int{100, 200, 300})

	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Println("last array data : ", array)

}
