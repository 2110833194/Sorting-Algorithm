package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
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

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ { //因为minIndex起始是i，所有j比较从i+1开始
			if arr[j] < arr[minIndex] {
				minIndex = j //记录最小数的下标
			}
		}
		if minIndex != i { //因为采取的异或交换，若两个值相等，会变成0，所有进行判断
			swap(arr, i, minIndex)
		}
	}
}

func main() {
	//arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 4.0 //设置数组的数据大小级别比如5就是10^5级别
	for i:=0.0;i<math.Pow(10,powNumer);i++{
		arr = append(arr,rand.Intn(100000))//随机生成数加入到数组中
	}	
	start := time.Now().UnixNano()
	selectionSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("普通选择排序对10^%d级别的数据排序需要%fs\n",int(powNumer),float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)
}
