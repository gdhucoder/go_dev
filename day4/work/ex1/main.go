package main

import "fmt"

// 冒泡排序
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		flag := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				swap(&a[j], &a[j+1])
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

// 实现一个选择排序
func SelectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		maxIdx := 0
		for j := 1; j < len(a)-i; j++ {
			if a[j] > a[maxIdx] {
				maxIdx = j
			}
		}
		swap(&a[len(a)-i-1], &a[maxIdx])
	}
}

// 实现一个插入排序
func InsertionSort(a []int) {
	// pick a[i], insert into sorted subarray a[0...(i-1)]
	for i := 1; i < len(a); i++ {
		fmt.Println(a)
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				swap(&a[i], &a[j])
			}
		}
	}
}

func QSort(a []int) {
	QuickSort(a, 0, len(a)-1)
}

// 实现一个快排
func QuickSort(a []int, low int, hi int) {

	if low >= hi {
		return
	}

	p := partition(a, low, hi)

	QuickSort(a, low, p-1)
	QuickSort(a, p+1, hi)

}

func partition(a []int, low int, hi int) int {

	// 5213 4
	// 4213 5
	pivot := a[hi] // 最后一个元素
	i := low - 1
	for j := low; j <= hi-1; j++ {
		if a[j] <= pivot {
			i++
			swap(&a[j], &a[i])
		}
	}
	swap(&a[i+1], &a[hi])
	return i + 1
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a := []int{5, 3, 2, 1, 4}
	// a = []int{1, 2, 3, 4, 5}
	// BubbleSort(a)
	// SelectionSort(a)
	// InsertionSort(a)
	QSort(a)
	fmt.Println(a)
}
