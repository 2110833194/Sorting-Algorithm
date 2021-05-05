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
	//arr := []int{3,44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 4.0 //设置数组的数据大小级别比如5就是10^5级别
	for i:=0.0;i<math.Pow(10,powNumer);i++{
		arr = append(arr,rand.Intn(100000))//随机生成数加入到数组中
	}	
	start := time.Now().UnixNano()
	bubbleSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("普通冒泡排序对10^%d级别的数据排序需要%fs\n",int(powNumer),float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)
}
