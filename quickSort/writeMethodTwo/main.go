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
//双指针法
func quickSort(arr []int){

	quickSortRescursion(arr,0,len(arr)-1)

}

func quickSortRescursion(arr []int,start,end int){
	if start >= end {
		return 
	}
	//求分区中轴，把数组分区
	middle := partition(arr,start,end)
	//向左分区递归，对左分区进行排序
	quickSortRescursion(arr,start,middle-1)
	//向右分区递归，对右分区进行排序
	quickSortRescursion(arr,middle+1,end)
} 

func partition(arr []int,start,end int)int{
	//基数
	// pivot := arr[start]
	index := rand.Intn(end+1)%(end-start)+start
	if index != start{
		swap(arr,start,index)//为了随机一点
	}
	pivot := arr[start]
    left := start+1
	right := end

	for left < right{
		for left < right && arr[left] <= pivot{
			left++
		}
		for left < right && arr[right] >= pivot{
			right--
		} 

		if left < right{
			swap(arr,left,right)
			left++
			right--
		}
	}
      // 如果 left 和 right 相等，单独比较 arr[right] 和 pivot
    if left == right && arr[right] > pivot{
		right--
	}
	// 将基数和轴交换
    if right != start{
	swap(arr,right,start)
	}
	return right
}

func main() {
	// arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 7.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	quickSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("双指针快速排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	// quickSort(arr)
	// fmt.Println(arr)
}