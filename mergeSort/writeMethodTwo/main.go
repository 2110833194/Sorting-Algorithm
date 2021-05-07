package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
//优化归并排序
func mergeSort(arr []int){
	if len(arr) <= 1{
		return
	}
	//辅助数组
	res := make([]int,len(arr))
	mergeRescurison(arr,0,len(arr)-1,res)
}

func mergeRescurison(arr []int,start,end int,res []int){
	if start >= end{
		return 
	}
	middle := (start+end)/2
    // 拆分左边区域，并将归并排序的结果保存到 result 的 [start, middle] 区间
	mergeRescurison(arr,start,middle,res)
	// 拆分右边区域，并将归并排序的结果保存到 result 的 [middle + 1, end] 区间
    mergeRescurison(arr,middle+1,end,res)
    // 合并左右区域到 result 的 [start, end] 区间
    merge(arr,start,end,res)
}

// 将 result 的 [start, middle] 和 [middle + 1, end] 区间合并
func merge(arr []int,start,end int,res []int){
	middle := (start+end)/2
	// 数组 1 的首尾位置
	start1,end1 := start,middle
	// 数组 2 的首尾位置
	start2,end2 := middle+1,end
    // 用来遍历数组的指针
	index1 := start1
	index2 := start2
    // 结果数组的指针
	resIndex := start
	for index1 <= end1 && index2 <= end2{
		if arr[index1] <= arr[index2]{
			res[resIndex] = arr[index1] 
			index1++
			resIndex++
		}else{
			res[resIndex] = arr[index2] 
			index2++
			resIndex++
		}
	} 
	//把剩余的数字移到数组里面去
	for index1 <= end1{
		res[resIndex] = arr[index1] 
		index1++
		resIndex++
	}
	
	for index2 <= end2{
		res[index1+index2-start2] = arr[index2] 
		index2++
		resIndex++
	}
	// 将 result 操作区间的数字拷贝到 arr 数组中，以便下次比较
    for i := start; i <= end; i++{
        arr[i] = res[i]
	}

}


func main() {
	// arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	// mergeSort(arr)
	// fmt.Println(arr)
	var arr []int
	powNumer := 7.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	mergeSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("一次优化归并排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))

}