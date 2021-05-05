package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func shellSort(arr []int){
	//创建 间隔序列（增量序列）
	for gap:=len(arr)/2;gap>0;gap/=2{
		for i:=gap;i<len(arr);i++{
			currentNumber := arr[i]//记录当前需要插入的数
			j := i-gap//记录当前下标的前一个下标
			for j>=0 && arr[j] > currentNumber{
				arr[j+gap] = arr[j] //后移元素
				j -= gap //更新下标
			}
			arr[j+gap] = currentNumber
		}

	}
}

func main() {
	//arr := []int{84,83,88,87,61,50,70,60,80,99}
	var arr []int
	powNumer := 6.0 //设置数组的数据大小级别比如5就是10^5级别
	for i:=0.0;i<math.Pow(10,powNumer);i++{
		arr = append(arr,rand.Intn(100000))//随机生成数加入到数组中
	}	
	start := time.Now().UnixNano()
	shellSort(arr)
	end := time.Now().UnixNano()
	fmt.Printf("优化了的希尔排序对10^%d级别的数据排序需要%fs\n",int(powNumer),float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)

}
