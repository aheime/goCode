package main

import "fmt"

func main() {
	str := "aheime"
	slice := str[:2]
	for i := range str {
		println(str[i])
		fmt.Printf("str: %v, %p\n", str, str[i])
	}

	fmt.Printf("str: %v, %p\n", str, &str)
	fmt.Printf("slice: %v, %p\n", slice, &slice)
	slice1 := []rune(str)
	fmt.Printf("slice1: %v, %d, %p\n", slice1, cap(slice1), &slice1[0])
	slice1[0] = '北'
	slice1[1] = '京'
	fmt.Printf("slice1: %v, %d, %p\n", slice1, cap(slice1), &slice1[0])
	slice1 = append(slice1, '啊')
	fmt.Printf("slice1: %v, %d, %p\n", slice1, cap(slice1), &slice1[0])

	fmt.Println(string(slice1))

}
