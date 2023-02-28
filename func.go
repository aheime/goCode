package main

import "fmt"

type myFunc func() int

// 闭包
func a() myFunc {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {

	var c = a()
	w := c()
	fmt.Println(w)
	c()
	c()
	a()
	var c2 = a()
	c2()
	c2()
}
