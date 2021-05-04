package main

import "fmt"

func insertSort(arr []int) {
	// 从第二个数开始，往前插入数字
	for i := 1; i < len(arr); i++ {
		currentNumber := arr[i]
		j := i - 1
		// 寻找插入位置的过程中，不断地将比 currentNumber 大的数字向后挪
		for j >= 0 && currentNumber < arr[j] {
			arr[j+1] = arr[j]
			//更新下标
			j--
		}
		// 两种情况会跳出循环：1. 遇到一个小于或等于 currentNumber 的数字，跳出循环，currentNumber 就坐到它后面。
		// 2. 已经走到数列头部，仍然没有遇到小于或等于 currentNumber 的数字，也会跳出循环，此时 j 等于 -1，currentNumber 就坐到数列头部。
		arr[j+1] = currentNumber

	}

}

func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	insertSort(arr)
	fmt.Println(arr)
}
