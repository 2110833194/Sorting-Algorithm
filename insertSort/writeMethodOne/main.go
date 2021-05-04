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

//交换法插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j >= 1 && arr[j-1] > arr[j] {
			swap(arr, j-1, j)
			j--
		}
	}

}

func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	insertSort(arr)
	fmt.Println(arr)
}
