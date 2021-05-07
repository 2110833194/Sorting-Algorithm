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
func selectionSort(arr []int) {
	// i 只需要遍历一半
	for i := 0; i < len(arr)/2; i++ {
		maxIndex := i
		minIndex := i
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		//如果 minIndex 和 maxIndex 都相等，那么他们必定都等于 i，且后面的所有数字都与 arr[i] 相等，此时已经排序完成
		if minIndex == maxIndex {
			break
		}
		if arr[i] != arr[minIndex] { //因为采取的异或交换，若两个值相等，会变成0，所有进行判断
			swap(arr, i, minIndex)
		}
		if maxIndex == i { // 如果最大值的下标刚好是 i，由于 arr[i] 和 arr[minIndex] 已经交换了，所以这里要更新 maxIndex 的值。
			maxIndex = minIndex
		}
		lastIndex := len(arr) - i - 1 //末尾下标
		if arr[maxIndex] != arr[lastIndex] {
			swap(arr, maxIndex, lastIndex) //最大值交换到末尾
		}
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
	selectionSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("二元选择排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)
}
