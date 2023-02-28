package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 向标准输出写入内容
	/*	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

		fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("打开文件出错，err:", err)
			return
		}
		name := "枯藤"
		// 向打开的文件句柄中写入内容
		fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)*/

	//s1 := fmt.Sprint("枯藤")
	//name := "枯藤"
	//age := 18
	//s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	//s3 := fmt.Sprintln("枯藤")
	//fmt.Println(s1, s2, s3)
	//
	//err := fmt.Errorf("这是一个错误")
	//fmt.Println(err)

	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	//bufioDemo()
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
