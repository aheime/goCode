package main

import "fmt"

type errHandleFunc func(err interface{})

func try(func1 func(), func2 errHandleFunc) {
	defer func() {
		if err := recover(); err != nil {
			func2(err)
		}
	}()

	func1()
}

func main() {

	try(func() {
		var num interface{} = 1
		v, ok := num.(int)
		fmt.Println(v, ok)
		fmt.Printf("%T", num)
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})

}
