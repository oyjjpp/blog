//算法
package models

import "fmt"

//@link 排序算法 https://juejin.im/post/5a08cc646fb9a045030f9174

/**
 * 插入排序 {5, 2, 0, 1, 3, 1, 4}
 * 从一组元素中取一个元素为有序元素组，然后在剩下的元素中每次取一个元素向有序的元素组插
 * 时间复杂度O(n^2)
 * 空间复杂度O(1)
 * 稳定性：稳定
 */
func InsertSort(item []int) []int {
	var temp int
	var j int

	for i := 1; i < len(item); i++ {
		//保存第一个值
		temp = item[i]
		j = i - 1

		for ; j >= 0 && item[j] > temp; j-- {
			item[j+1] = item[j]
		}
		item[j+1] = temp
	}
	return item
}

/**
 * @desc 希尔排序
 * 时间复杂度O(n^1.3)
 * 空间复杂度O(1)
 * 稳定性：不稳定
 */
func ShellSort(item []int) []int {
	var temp int
	var j int

	n := len(item)
	for d := n / 2; d > 0; d = d / 2 {
		for x := 0; x < d; x++ {
			for i := x + d; i < n; i = i + d {
				temp = item[i]
				j = i - d
				for ; j >= 0 && item[j] > temp; j = j - d {
					item[j+d] = item[j]
				}
				item[j+d] = temp
			}
		}
	}
	return item
}

/**
 * 简单选择排序
 * 循环查找“最小”元素放在首位
 * 时间复杂度O(n^2)
 * 空间复杂度O(1)
 * 稳定性 : 不稳定
 */
func SelectSort(item []int) []int {
	var j int
	var temp int
	var position int

	for i, n := 0, len(item); i < n; i++ {
		j = i + 1
		temp = item[i]
		position = i
		for ; j < n; j++ {
			if item[j] < temp {
				temp = item[j]
				position = j
			}
		}
		item[position] = item[i]
		item[i] = temp
	}

	return item
}

/**
 * @desc 创建最大堆
 * @param slice item 元素组
 * @param heapSize int 需要创建最大堆的大小
 * @param index int 当前需要创建最大堆的位置
 */
func maxHeapify(item []int, heapSize, index int) {
	left := index*2 + 1
	right := left + 1

	largest := index

	if left < heapSize && item[index] < item[left] {
		largest = left
	}

	if right < heapSize && item[largest] < item[right] {
		largest = right
	}

	if largest != index {
		temp := item[index]
		item[index] = item[largest]
		item[largest] = temp
		maxHeapify(item, heapSize, largest)
	}
}

/**
 * @desc 堆排序
 * @param 时间复杂度 O(nlogn)
 */
func HeadSort(item []int) []int {
	n := len(item)
	startIndex := (n - 1 - 1) / 2

	for i := startIndex; i >= 0; i-- {
		maxHeapify(item, n, i)
	}

	var temp int
	for i := n - 1; i > 0; i-- {
		temp = item[0]
		item[0] = item[i]
		item[i] = temp
		maxHeapify(item, i, 0)
	}

	return item
}

/**
 * @desc 冒泡排序
 * 时间复杂度O(n^2)
 * 空间复杂度O(1)
 * 稳定性:稳定
 */
func BubbleSort(item []int) []int {
	n := len(item)

	for i := 0; i < n-1; i++ {
		for j := n - 1 - 1; j >= i; j-- {
			if item[j+1] < item[j] {
				temp := item[j]
				item[j] = item[j+1]
				item[j+1] = temp
			}
		}
	}
	return item
}

/**
 * @desc 快速排序
 * 时间复杂度O(nlogn)
 */
func QuikcSort(item []int) []int {
	quickSort(item, 0, len(item)-1)
	return item
}

//交换
func swap(item []int, i, j int) {
	temp := item[i]
	item[i] = item[j]
	item[j] = temp
}

//快速排序
func quickSort(item []int, start, end int) {
	if start < end {
		pivot := item[start]
		left := start
		right := end

		for left != right {
			for item[right] >= pivot && left < right {
				right--
			}

			for item[left] <= pivot && left < right {
				left++
			}
			println(left, right)
			swap(item, left, right)
			fmt.Printf("%v\n", item)
		}

		item[start] = item[left]
		item[left] = pivot
		quickSort(item, start, left-1)
		quickSort(item, left+1, end)
	}
}

/**
 * @desc 归并排序
 * 时间复杂度 O(nlog2n)
 * 空间复杂度 O(n) + O(log2n)
 * 稳定性：稳定
 */
func MergeSort(item []int) []int {
	mergeSort(item, 0, len(item)-1)
	return item
}

func mergeSort(item []int, left, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSort(item, left, center)
		mergeSort(item, center+1, right)
		merge(item, left, center, right)
	}
}

func merge(item []int, left, center, right int) {
	var tempItem []int

	mid := center + 1

	//记录中间数组的索引
	third := left

	//复制是用到的索引
	temp := left

	for left <= center && mid <= right {
		if item[left] <= item[mid] {
			third++
			left++
			tempItem[third] = item[left]
		} else {
			third++
			mid++
			tempItem[third] = item[mid]
		}
	}

	for mid <= right {
		third++
		mid++
		tempItem[third] = item[mid]
	}

	for left <= center {
		third++
		left++
		tempItem[third] = item[left]
	}

	for temp <= right {
		item[temp] = tempItem[temp]
	}
}

func RadixSort(item []int) []int {
	max := item[0]
	itemLen := len(item)

	for i := 1; i < itemLen; i++ {
		if item[i] > max {
			max = item[i]
		}
	}

	time := 0 //数组最大值位数
	for max > 0 {
		max = max / 10
		time++
	}

	return item
}
