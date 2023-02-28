package main

import (
	"errors"
	"fmt"
)

//1. 关键字 defer 用于注册延迟调用。
//2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
//3. 多个defer语句，按先进后出的方式执行。
//4. defer语句中的变量，在defer声明时就决定了。

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)
}

// 如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值
func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func main() {

	//arr := [5]int{}
	//先进后出
	//for i := range arr {
	//	defer fmt.Println(i)
	//}

	//函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4
	//for i := range arr {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}

	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		//defer t.Close()
		defer fmt.Println(t)
		t2 := t
		defer t2.Close()
		defer Close(t) //defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行
	}

	//test()

	foo(2, 0)

	fmt.Println("defer")
}
