package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//var buf []byte
	//os.Stdin.Read(buf[:])
	//os.Stdout.WriteString(string(buf[:]))

	// 打开文件
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
