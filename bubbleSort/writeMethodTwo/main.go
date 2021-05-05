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

//经过优化的写法：使用一个变量记录当前轮次的比较是否发生过交换，如果没有发生交换表示已经有序，不再继续排序；
func bubbleSort(arr []int) {
	swapped := true //记录是否发生交换
	for i := 0; i < len(arr)-1; i++ {
		if !swapped { //若上一轮未发生交换，就退出排序
			break
		}
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				swapped = true //发生交换
			}
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
	bubbleSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("一次优化的冒泡排序插入排序对10^%d级别的数据排序需要%fs\n",int(powNumer),float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)
}
