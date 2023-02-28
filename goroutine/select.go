package main

import (
	"fmt"
)

func main() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()
	go func() {
		string_chan <- "hello"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	default:
		// 如果两个通道都没有可用的数据，则执行这里的语句
		fmt.Println("no message received")
	}

	fmt.Println("main结束")
}
