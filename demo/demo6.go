package main

import "fmt"

type Test struct {
	Index int
	Num   int
}

func main() {
	var t []*Test
	t = append(t, &Test{Index: 1, Num: 1})
	t = append(t, &Test{Index: 2, Num: 2})

	for _, v := range t {
		v.Num += 100
	}

	for _, v := range t {
		fmt.Printf("%#v", v)
	}

	a := ^8
	fmt.Println(a)

	for i := 1; i <= 10; i++ {
		if i&1 == 1 {
			fmt.Println(i)
		}
	}

}
