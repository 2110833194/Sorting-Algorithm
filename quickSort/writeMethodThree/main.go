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
//挖坑法，随机基准数
func quickSort(arr []int){
     quickSortRescursion(arr,0,len(arr)-1)
}

func quickSortRescursion(arr []int,start,end int){
    if start >= end{
		return //递归终止条件
	} 
	middle := partition(arr,start,end)
    quickSortRescursion(arr,start,middle-1)
	quickSortRescursion(arr,middle+1,end)
}
		
func partition(arr []int,start,end int)int{
    // index := (start+end) >> 1
	//随机选取基准数
	index := rand.Intn(end+1)%(end-start)+start
	if index != start{
		swap(arr,start,index)//为了随机一点
	}
	pivot := arr[start]
	left := start
    right := end
	for left<right{
		for left < right && arr[right] > pivot{
			right--
		} 
		if left < right{
			arr[left] = arr[right]
		    left ++
		}
		for left < right && arr[left] <= pivot{
			left++
		}
		if left < right{
            arr[right] = arr[left]
		    right --
	    }
	}
        arr[left] = pivot
		return left 
}



func main() {
	// arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 8.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	quickSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("挖坑法快速排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	// quickSort(arr)
	// fmt.Println(arr)
}