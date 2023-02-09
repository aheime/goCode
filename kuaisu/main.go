package main

import "fmt"

func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	temp := 0
	// for 循环的目标是将比 pivot 小的数放到左边，比 pivot 大的数放到右边
	for l < r {
		// 从 pivot 的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		// 从 pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		// 交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		// 优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func quickSort(arr []int) []int {
	n := len(arr)
	fmt.Println(arr)

	//如果n为1，即数组只要一个元素，不需要排序，返回即可
	if n < 2 {
		return arr
	} else {
		middle := arr[0] //获取比较的参考值（基准值）
		var low []int    //小于基准值的数据组成的切片
		var high []int   //大于基准值的数据组成的切片
		for i := 1; i < n; i++ {
			if arr[i] < middle {
				low = append(low, arr[i]) //小于基准值的数据归类
			} else {
				high = append(high, arr[i]) //大于基准值的数据归类
			}
		}

		/*
			以下为组合成的切片：quickSort（low）+middle+quickSort(high)
			返回值为：   递归小于部分 + 基准值  +  递归大于部分
		*/
		lowSlice := quickSort(low)
		highSlice := quickSort(high)
		res := append(lowSlice, middle)
		for _, data := range highSlice {
			res = append(res, data)
		}
		fmt.Println("res=>", res)

		return res
	}

}

func main() {
	//arr := [9]int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	////调用快速排序
	//QuickSort(0, len(arr)-1, &arr)
	//fmt.Println(arr)

	s := []int{22, 4, 7, 8, 8, 6, 3, 10, 9, 88, 66, 22}
	fmt.Println(quickSort(s))
}
