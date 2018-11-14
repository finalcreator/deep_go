package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Quick Sort in Golang
func main() {

	slice := generateAndShuffleSlice4(10)
	fmt.Println("\n--- Unsorted --- \n", slice)
	// quickSortV1(slice, 0, len(slice)-1)
	// quickSortV2(slice)
	// quickSortV3(slice, 0, len(slice)-1)
	//quickSortV4(slice)
	fmt.Println("\n--- Sorted ---\n", slice)
}

func quickSortV1(values []int, left int, right int) {
	if left < right {

		// 设置分水岭
		temp := values[left]

		// 设置哨兵
		i, j := left, right
		for {
			// 从右向左找，找到第一个比分水岭小的数
			for values[j] >= temp && i < j {
				j--
			}

			// 从左向右找，找到第一个比分水岭大的数
			for values[i] <= temp && i < j {
				i++
			}

			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}

		// 将分水岭移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		quickSortV1(values, left, i-1)
		quickSortV1(values, i+1, right)
	}
}

func quickSortV2(values []int) {
	fmt.Println("-----------------------------------")
	fmt.Printf("\nSort Silice %v\n", values)
	length := len(values)
	if length <= 1 {
		return
	}
	mid, i := values[0], 1    // 取第一个元素作为分水岭，i下标初始为1，即分水岭右侧的第一个元素的下标
	head, tail := 0, length-1 // 头尾的下标

	// 如果头和尾没有相遇，就会一直触发交换
	for head < tail {
		fmt.Printf("Before Swap {head:%d,tail:%d,i:%d,mid:%d} values:%v\n", head, tail, i, mid, values)
		if values[i] > mid {
			// 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			fmt.Printf("(values[i] > mid) values[i=%d]=%d <-> values[tail=%d]=%d\n", i, values[i], tail, values[tail])
			values[i], values[tail] = values[tail], values[i]
			tail-- // 尾下标左移一位
		} else {
			// 如果分水岭右侧的元素小于等于分水岭，就将分水岭右移一位
			fmt.Printf("(values[i] <= mid) values[i=%d]=%d <-> values[head=%d]=%d\n", i, values[i], head, values[head])
			values[i], values[head] = values[head], values[i]
			head++ // 头下标右移一位
			i++    // i下标右移一位
		}
		fmt.Printf("After Swap {head:%d,tail:%d,i:%d,mid:%d} values:%v\n\n", head, tail, i, mid, values)
	}

	// 分水岭左右的元素递归做同样处理
	quickSortV2(values[:head])
	quickSortV2(values[head+1:])
}

func quickSortV3(a []int, left, right int) {
	if left >= right {
		return
	}
	val := a[left] //基准值
	k := left      //当前基准值位置

	//逐个查找小于基准值的值
	for i := left + 1; i <= right; i++ {
		if a[i] < val { //找到a[i]小于基准值
			a[k] = a[i]   //当前基准值所在位置变更为a[i]
			a[i] = a[k+1] //a[i]变更为 a[k+1] ,因为a[k+1] > 基准值
			k++           //变更基准值位置,暂时不用写入基准值,只需记录位置
		}
	}
	a[k] = val //最后写入基准值
	quickSortV3(a, left, k-1)
	quickSortV3(a, k+1, right)
}

func quickSortV4(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	quickSortV4(a[:left])
	quickSortV4(a[left+1:])

}

func generateAndShuffleSlice1(size int) []int {

	list := rand.Perm(size)
	for i := range list {
		list[i]++
	}
	return list

}

func generateAndShuffleSlice2(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i + 1
	}
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s

}

func generateAndShuffleSlice3(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i + 1
	}
	for i := range slice {
		j := rand.Intn(size)
		//fmt.Println(j)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice

}

func generateAndShuffleSlice4(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i + 1
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice

}

// Generates a slice of size, size filled with random numbers
// func generateSlice(size int) []int {

// 	slice := make([]int, size, size)
// 	rand.Seed(time.Now().UnixNano())
// 	intSet := myset.NewSet()
// 	for i := 0; i < size; i++ {
// 		// slice[i] = rand.Intn(999) - rand.Intn(999)
// 		// slice[i] = rand.Intn(20)

// 		intSet.Add(rand.Intn(999))
// 	}

// 	for k := range intSet.GetInnerMap() {
// 		fmt.Println(k)
// 		slice = append(slice, k)
// 	}
// 	return slice
// }
