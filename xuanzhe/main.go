package main

import "fmt"

func main() {

	var arrs = [...]int{3, 10, 4, 1, 8}

	leng := len(arrs)

	for j := 0; j < leng-1; j++ {
		maxIndex := j
		max := arrs[j]
		for i := j + 1; i < leng; i++ {
			if max < arrs[i] {
				max = arrs[i]
				maxIndex = i
			}
		}

		if maxIndex != j {
			arrs[j], arrs[maxIndex] = arrs[maxIndex], arrs[j]
		}

		fmt.Printf("第%d次arr=%v\n", j+1, arrs)
	}

}
