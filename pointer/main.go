package main

import "fmt"

func main() {
	var a *int
	var b = 10
	a = &b
	fmt.Printf("a指针变量：%v，类型为%T\n", a, a)
	fmt.Printf("a的地址：%v\n", &a)
	fmt.Printf("b的地址：%v，类型为%T\n", &b, b)
	fmt.Printf("a指针变量的值：%v\n", *a)

	var arr1 [5]int
	pp := new([5]int)
	pp = &arr1
	fmt.Printf("%p\n", &arr1)
	fmt.Println(&arr1)
	fmt.Println(*pp)
	pp[1] = 10
	fmt.Println(&arr1[0])
	fmt.Println(&arr1[1])
	fmt.Println(arr1)
}
