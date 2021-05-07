package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)


func swap(nums []int, i, j int) {
	nums[i] = nums[i] ^ nums[j]
	nums[j] = nums[j] ^ nums[i]
	nums[i] = nums[i] ^ nums[j]

}

func quickSort(arr []int){

	quickSortRescursion(arr,0,len(arr)-1)

}

func quickSortRescursion(arr []int,start,end int){
	if start >= end {
		return 
	}
	//求分区中点，把数组分区
	middle := partition(arr,start,end)
	//向左分区递归，对左分区进行排序
	quickSortRescursion(arr,start,middle-1)
	//向右分区递归，对右分区进行排序
	quickSortRescursion(arr,middle+1,end)
}

func partition(arr []int,start,end int)int{
	//基数
	pivot := arr[start] 
	//左下标，从第二个数开始
	left := start+1
	//右下标
	right := end
	for left < right{
		for left < right && arr[left] <= pivot{
			left++
		}
		//如果剩余的数组都比基数小，则 left 会加到 right 才停止，
		//这时不应该发生交换。因为 right 已经指向了最后一个比基数小的值。
		if left != right{
			swap(arr,left,right)
			right--
		}
		
	}

	//1.剩余数组中只有最后一个数比基数大的情况
	//2.left 和 right 区间内只有一个值，则初始状态下， left == right，所以 while (left < right) 根本不会进入，
	//所以此时我们单独比较这个值和基数的大小关系
	//3.是剩余数组中每个数都比基数大，此时 right 会持续减小，直到和 left 相等退出循环，
	//此时 left 所在位置的值还没有和 pivot 进行比较，所以我们单独比较 left 所在位置的值和基数的大小关系
	if left == right && arr[right] > pivot{
		right --
	}

	if start != right{
		swap(arr,start,right)
	}
	return right
}


func main() {
	// arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 6.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	quickSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("简单分区快速排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	// quickSort(arr)
	// fmt.Println(arr)
}