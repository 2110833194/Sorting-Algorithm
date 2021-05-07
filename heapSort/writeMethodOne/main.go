package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func swap(arr []int, i, j int) {
	//交换元素（异或）
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[j] ^ arr[i]
	arr[i] = arr[i] ^ arr[j]
}

//大顶堆
func maxHeapSort(arr []int) {
	//如何用数列构建出一个大顶堆；
	//取出堆顶的数字后，如何将剩余的数字调整成新的大顶堆。
	// 构建初始大顶堆
	buildMaxHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)       //交换元素，使得当前堆最大的元素交换到堆的末尾
		maxHeapify(arr, 0, i) //从0第一个非叶子节点开始调整堆,用i是因为初建堆已经把堆的大小-1了
	}

}

func buildMaxHeap(arr []int) {
	// 从最后一个非叶子结点开始调整大顶堆，最后一个非叶子结点的下标就是 arr.length / 2-1
	for i := len(arr)/2 - 1; i >= 0; i-- { //i--是为了不断从前面的非叶子节点调整堆
		maxHeapify(arr, i, len(arr))
	}

}

func maxHeapify(arr []int, index, heapSize int) {
	// 左子结点下标
	l := index<<1 + 1 //当前调整堆节点的左子节点
	// 右子结点下标
	r := l + 1 //当前调整堆节点的右子节点
	// 记录根结点、左子树结点、右子树结点三者中的最大值下标
	largest := index // 记录最大数的下标
	// 与左子树结点比较
	if l < heapSize && arr[l] > arr[largest] {
		largest = l
	}
	// 与右子树结点比较
	if r < heapSize && arr[r] > arr[largest] {
		largest = r
	}
	//递归向下调整
	if largest != index { //若不相等就说明
		swap(arr, largest, index)          //交换元素，使其符合大顶堆的性质
		maxHeapify(arr, largest, heapSize) //递归，把最大元素和他的子树节点相比较，找出最大元素，
	}

}

func main() {
	//arr := []int{3, 44, 38, 5, 47, 15, 36, 34}
	var arr []int
	powNumer := 7.0 //设置数组的数据大小级别比如5就是10^5级别
	for i := 0.0; i < math.Pow(10, powNumer); i++ {
		arr = append(arr, rand.Intn(100000)) //随机生成数加入到数组中
	}
	start := time.Now().UnixNano()
	maxHeapSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("大顶堆排序对10^%d级别的数据排序需要%fs\n", int(powNumer), float64(end-start)/(math.Pow(10, 9)))
	// maxHeapSort(arr)
	// fmt.Println(arr)
}
