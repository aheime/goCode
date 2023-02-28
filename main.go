package main

import (
	"fmt"
)

func main() {
	var a = 'a'
	var b byte = 'a'

	fmt.Printf("%T, %v\n", a, a) // int32, 97
	fmt.Printf("%T, %v\n", b, b) // uint8, 97

	b1 := byte('a')             // 字符
	b2 := []byte("A")           // 字符串
	b3 := []byte{'a', 'b', 'c'} // 字符串
	b4 := []byte("你好")
	b5 := []rune("你好m")
	var b6 = '\u4f60'
	fmt.Printf("b1 = %c\n", b1)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Printf("b3 = %s\n", b3)
	fmt.Printf("b4 = %v\n", b4)
	fmt.Printf("b5 = %v\n", b5)
	fmt.Printf("b6 = %v\n", b6)

	for _, v := range b5 {
		fmt.Printf("unicode: %c %d %X %U %v \n", v, v, v, v, v)
	}
}
