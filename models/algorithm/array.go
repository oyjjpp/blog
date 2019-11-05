//数组常见相关算法
package algorithm

//数组中求最大的数，并且是其他数的两倍
//解题思路：暴力破解，进行两层循环比较
func dominantIndex(nums []int) int {
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
func dominantIndexV2(nums []int) int{
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
func dominantIndexV3(nums []int) int{
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
