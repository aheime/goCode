package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包

	fmt.Println(a)

	arr := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := arr[1:3:4]
	fmt.Printf("s:len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	fmt.Printf("arr[1]:地址: %p\n", &arr[1])
	s = append(s, 100)
	fmt.Printf("s:len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	s[0] = 11 //会改变arr
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("s: %v\n", s)
	s = append(s, 200) //超出 s.cap 限制,底层生成新的arr
	fmt.Printf("s:len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	s[0] = 22 //不会改变arr
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("s: %v\n", s)

}
