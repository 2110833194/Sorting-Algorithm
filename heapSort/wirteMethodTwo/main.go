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

//小顶堆
func minHeapSort(nums []int) {
	buildMinheap(nums) //初建顶堆
	//需要调整堆k-1次（初建堆已经完成一次）

	for i := len(nums) - 1; i > 0; i-- {
		swap(nums, 0, i)
		minHeapify(nums, 0, i)
	}

}

func buildMinheap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- { //建初堆
		minHeapify(nums, i, len(nums))
	}

}

func minHeapify(nums []int, index, heapSize int) {
	//调整堆
	l := (index << 1) + 1
	r := l + 1
	smallest := index

	if l < heapSize && nums[l] < nums[smallest] {
		smallest = l
	}

	if r < heapSize && nums[r] < nums[smallest] {
		smallest = r
	}

	if smallest != index {
		//交换,使其满足最小堆性质
		swap(nums, smallest, index)
		//向下递归调整
		minHeapify(nums, smallest, heapSize)
	}

}

func main() {
	// arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 7.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	minHeapSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("小顶堆排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	// minHeapSort(arr)
	// fmt.Println(arr)
}
