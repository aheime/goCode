package main

import (
	"flag"
	"fmt"
	"time"
)

//Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。

func main() {

	//if len(os.Args) > 0 {
	//	for i, v := range os.Args {
	//		fmt.Printf("第%d个参数是：%v", i, v)
	//	}
	//}

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
