//数组常见相关算法
package algorithm

import (
	"blog/models"
)

//长度最小的子数组
func MinSubArrayLen(s int, nums []int) int {
	k:=0
	for i:=0;i<nums;i++{

	}
	return k
}


//最大连续1的个数
func FindMaxConsecutiveOnes(nums []int) int {
	var rs ,temp int
	for i:= 0; i< len(nums);i++{
		if nums[i] == 1 {
			temp++
		}else{
			if (temp > rs){
				rs = temp
			}
			temp = 0
		}
	}
	//考虑最后一个也为1的情况
	if temp >= rs {
		rs = temp
	}
	return rs
}


//移除元素
//给定一个数组和一个值，原地删除该值的所有实例并返回新的长度
//解题思路：使用两个指针，一个用于迭代，一个指针总是指向下一次添加的位置
func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

//两数之和
//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//解题思路：循环相加 碰到相等的则返回
func TwoSum(numbers []int, target int) []int {
	rs := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			temp := numbers[i] + numbers[j]
			if temp == target {
				rs[0] = i + 1
				rs[1] = j + 1
			}
		}
	}
	return rs
}

//数组拆分
//给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。
//解题思路：数组排序，连续两个为一组，即可计算出最大的和
func ArrayPairSum(nums []int) int {
    mem := models.QuikcSort(nums)
	rs := 0
	for i := 0; i < len(mem); i = i + 2 {
		rs = rs + mem[i]
	}
	return rs
}  
//反转数组中的元素
//双指针技巧使用场景
func ReverseArray(mem []int) []int {
	if len(mem) < 2 {
		return mem
	}
	i := 0
	n := len(mem) - 1
	for i < n {
		mem[i], mem[n] = mem[n], mem[i]
		i++
		n--
	}
	return mem
}

//杨辉三角
//1、col=0  元素为1
//2、col=row 元素为1 并不在为当前行余下元素赋值
func Generate(numRows int) [][]int {
	if numRows < 0 {
		return [][]int{}
	}

	//声明一个二维切片
	res := make([][]int, numRows)
	for index := 0; index < numRows; index++ {
		res[index] = make([]int, index+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numRows; j++ {
			if j == i {
				res[i][j] = 1
				break
			} else if j == 0 {
				res[i][j] = 1
			} else if i > 0 && j > 0 {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}

//螺旋矩阵
//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素
//定义上、下、左、右四个元素循环处理
func SpiralOrder(matrix [][]int) []int {
	//数组元素为空
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	//请求当前数组的列数、行数
	col, row := len(matrix), len(matrix[0])

	//定义 上线左右四个角的元素
	top, bottom, left, right := 0, col-1, 0, row-1
	count := 0

	nums := col * row
	res := make([]int, nums)

	for true {
		//从左到有遍历
		for i := left; i <= right; i++ {
			res[count] = matrix[top][i]
			count++
		}
		top++

		if left > right || top > bottom {
			break
		}
		//从上到下遍历
		for i := top; i <= bottom; i++ {
			res[count] = matrix[i][right]
			count++
		}
		right--
		if left > right || top > bottom {
			break
		}
		//从右到做遍历
		for i := right; i >= left; i-- {
			res[count] = matrix[bottom][i]
			count++
		}
		bottom--
		if left > right || top > bottom {
			break
		}
		//从下到上遍历
		for i := bottom; i >= top; i-- {
			res[count] = matrix[i][left]
			count++
		}
		left++
		if left > right || top > bottom {
			break
		}
	}
	return res
}

//以对角线方式遍历二维数字
//1、m-1,n+1；2、m+1,n-1
//注意索引溢出情况
func FindDiagonalOrder(matrix [][]int) []int {
	//数组元素为空
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	//请求当前数组的列数、行数
	col, row := len(matrix), len(matrix[0])

	//初始化元素总数、列数、行数
	nums, m, n := col*row, 0, 0
	res := make([]int, nums)
	flag := true

	//遍历
	for i := 0; i < nums; i++ {
		res[i] = matrix[m][n]

		if flag {
			//m-1,n+1【行-1，列+1】
			m = m - 1
			n = n + 1
		} else {
			//m+1,n-1【行+1，列-1】
			m = m + 1
			n = n - 1
		}

		//超过索引范围方式处理
		if m >= col {
			m = m - 1
			n = n + 2
			flag = true
		} else if n >= row {
			n = n - 1
			m = m + 2
			flag = false
		}

		if m < 0 {
			m = 0
			flag = false
		} else if n < 0 {
			n = 0
			flag = true
		}
	}
	return res
}

//加一
func PlusOne(digits []int) []int {
	num := len(digits)
	if num < 1 {
		return digits
	}

	for i := num - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	rs := make([]int, num+1)
	rs[0] = 1
	return rs
}

//数组中求最大的数，并且是其他数的两倍
//解题思路：暴力破解，进行两层循环比较
func DominantIndex(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return -1
	}

	if numsLen = 1{
		return 0
	}

	
	rs := false
	for i := 0; i < numsLen; i++ {
		for j := 0; j < numsLen; j++ {
			if j == i {
				continue
			}
			if nums[i] < (2 * nums[j]) {
				rs = true
				break
			}
		}

		if rs {
			rs = false
			continue
		}
		return i
	}
	return -1
}

//解题思路：遍历数组，找到最大值、次大值，最后进行比较
func DominantIndexV2(nums []int) int{
	numsLen := len(nums)
	if numsLen < 1 {
		return -1
	}
	index := 0
	maxNum := nums[0]
	secondNum := 0

	for i := 1; i < numsLen; i++ {
		if nums[i] > maxNum {
			index = i
			secondNum = maxNum
			maxNum = nums[i]

		} else if nums[i] > secondNum {
			secondNum = nums[i]
		}
	}

	if maxNum >= (2 * secondNum) {
		return index
	}
	return -1
}

//解题思路 假设第一个为最大值
func DominantIndexV3(nums []int) int{
	numsLen := len(nums)
	if numsLen < 1{
		return -1
	}
	index :=0
	maxNum := nums[0]

	for i:=1; i< numsLen;i++{
		if maxNum >= (2*nums[1]) {
			continue
		}else if nums[i] > (2*maxNum){
			index = i
			maxNum = nums[i]
		}else if nums[i] > maxNum{
			index =  -1
			maxNum = nums[i]
		}else{
			index = -1
		}
	}
	return index
}

//寻找数组的中心索引
//数组中心索引：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和
//特殊点：需要考虑第一个元素的情况
func PivotIndex(nums []int) int {
	//元素个数小于3则不存在中心索引
	if len(nums) < 1 {
		return -1
	}

	//存储中数据量
	var sum int
	//获取所有元素的和
	for _, v := range nums {
		sum = sum + v
	}
	//减去第一个元素
	sum = sum - nums[0]

	//存储左侧元素的总和
	var leftSum int
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if leftSum == sum {
			return i
		}
		leftSum = leftSum + nums[i]
		if (i + 1) < numsLen {
			sum = sum - nums[i+1]
		}
	}
	return -1
}
