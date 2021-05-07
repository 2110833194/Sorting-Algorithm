package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

//进一步优化的写法：除了使用变量记录当前轮次是否发生交换外，
//再使用一个变量记录上次发生交换的位置，下一轮排序时到达上次交换的位置就停止比较
func bubbleSort(arr []int) {
	swapped := true //记录是否发生交换
	indexOfLastUnsortedElement := len(arr) - 1
	// 上次发生交换的位置
	swappedIndex := -1
	for swapped {
		swapped = false
		for i := 0; i < indexOfLastUnsortedElement; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				swapped = true //发生交换
				swappedIndex = i
			}
		}
		indexOfLastUnsortedElement = swappedIndex
	}
}

func main() {
	//arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 4.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	bubbleSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("二次优化冒泡排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)
}
