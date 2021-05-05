package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)
//我们采用的增量序列为 D_m = N/2, D_k = D_{k+1} / 2,
//这个序列正是当年希尔发表此算法的论文时选用的序列，所以也被称之为希尔增量序列
func shellSort(arr []int){
		//生成间隔序列，在希尔排序中我们称之为增量序列
	for gap:=len(arr)/2;gap>0;gap/=2{ 
		//将一个数组转换成gap个数组（每一个数组各个元素之间的距离相差gap）
		for groupStartIndex:=0;groupStartIndex<gap;groupStartIndex++{	
		//插入排序(i:=groupStartIndex+gap就相当于插入排序的i:=1)
			for i:=groupStartIndex+gap;i<len(arr);i+=gap{
				currentNumber:=arr[i] 
				j:=i-gap
				for j >=0 && currentNumber < arr[j]{
					arr[j+gap] = arr[j]
					j -= gap 
				} 
				arr[j+gap] = currentNumber
			}	
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
	fmt.Printf("希尔排序对10^%d级别的数据排序需要%fs\n",int(powNumer),float64(end-start)/(math.Pow(10, 9)))
	//fmt.Println(arr)

}

