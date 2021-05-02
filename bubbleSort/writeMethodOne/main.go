package main

import (
	"fmt"
)

//交换函数的几种写法
/*
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

}
func swap(arr []int, i, j int) {
	arr[i] = arr[i] + arr[j]
	arr[j] = arr[i] - arr[j]
	arr[i] = arr[i] - arr[j]

}
*/

func swap(arr []int, i, j int) {
	//交换元素（异或）
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[j] ^ arr[i]
	arr[i] = arr[i] ^ arr[j]
}

//一边比较一边向后两两交换，将最大值 / 最小值冒泡到最后一位；
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}
func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	bubbleSort(arr)
	fmt.Println(arr)
}
