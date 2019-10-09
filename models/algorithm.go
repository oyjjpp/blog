//@link https://juejin.im/post/5a08cc646fb9a045030f9174
//算法
package models

/**
 * 插入排序 {5, 2, 0, 1, 3, 1, 4}
 * 时间复杂度O(n^2)
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
 * 简单排序
 * 时间复杂度O(n^2)
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
