package main

import "fmt"

// 在上面的示例中，quickSort函数实现了快速排序的逻辑，而partition函数则用于将数组划分为两部分，一部分小于等于基准值，一部分大于基准值。在主函数中，我们声明了一个整数类型的数组，调用quickSort函数对其进行排序，然后输出排序后的结果。
func quickSort(arr []int, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	fmt.Println("Before sorting:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sorting:", arr)
}
