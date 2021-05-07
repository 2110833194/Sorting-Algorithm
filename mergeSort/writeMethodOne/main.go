package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//归并排序
func mergeSort(arr []int){
	if len(arr) <= 1{
		return
	}
	//把结果存入res中
	res := mergeRescurison(arr,0,len(arr)-1)
	for i:=0;i<len(arr);i++{
		arr[i] = res[i]
	}

}
//拆分数组
func mergeRescurison(arr []int,start,end int)[]int{
	if start >= end{//递归终止条件
		return []int{arr[start]}
	}
	
	//左区间递归
	left := mergeRescurison(arr,start,(start+end)/2)
	//右区间递归
	right := mergeRescurison(arr,(start+end)/2+1,end)
	//将两个有序数组合并
	return merge(left,right)
}

//arr1,arr2两个有序数组合并到res中去
func merge(arr1,arr2 []int)[]int{
	index1,index2 := 0,0 
	res := make([]int,len(arr1)+len(arr2))
	for index1 < len(arr1) && index2 < len(arr2){
		if arr1[index1]<=arr2[index2]{
			res[index1+index2] = arr1[index1]
			index1++
		}else{
			res[index1+index2] = arr2[index2]
			index2++
		}
	}
	for index1 < len(arr1){
		res[index1+index2] = arr1[index1]
		index1++
	}

	for index2 < len(arr2){
		res[index1+index2] = arr2[index2]
		index2++
	}
	return res
}




func main() {
	//arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	//mergeSort(arr)
	//fmt.Println(arr)
	var arr []int
	powNumer := 7.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	mergeSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("经典归并排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))

}
