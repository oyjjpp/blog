//字符串类型常见算法问题
package algorithm

/*=====旋转字符串(start)=====*/
//暴力移位发 时间复杂度O(m*n) 空间复杂度O(1)
func LeftRotateString(data []byte, m int) []byte {
	for m > 0 {
		m--
		leftShiftOne(data)
	}
	return data
}

func leftShiftOne(data []byte) []byte {
	n := len(data)
	start := data[0]
	for i := 1; i < n; i++ {
		data[i-1] = data[i]
	}
	data[n-1] = start
	return data
}

//三步翻转发 时间复杂度O(n) 空间复杂度O(1)
func LeftRotateStringV2(data []byte, m int) {
	n := len(data)
	m = m % n
	reverseString(data, 0, m-1)
	reverseString(data, m, n-1)
	reverseString(data, 0, n-1)
}

func reverseString(data []byte, from, to int) {
	for from < to {
		str := data[from]
		data[from] = data[to]
		from++
		data[to] = str
		to--
	}
}

/*=====旋转字符串(end)=====*/
