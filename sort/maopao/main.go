package main

import "fmt"

func main() {

	var arr = [...]int{8, 10, 2, 0, 9}

	var lena = len(arr)
	for i := 0; i < lena-1; i++ {

		for j := 0; j < lena-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}

	}
	fmt.Println(arr)
}
