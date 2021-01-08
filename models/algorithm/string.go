//字符串类型常见算法问题
package algorithm

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//最长公共前缀
func LongestCommonPrefix(strs []string) string {
	strLen := len(strs)
	//空的字符串数组
	if strLen == 0 {
		return ""
	}
	//含有一个元素的字符串数组
	if strLen == 1 {
		return string(strs[0])
	}
	//假设前缀为第一个元素
	pre := string(strs[0])
	for i := 1; i < strLen; i++ {
		cur := strs[i]

		n := len(pre)
		if len(pre) >= len(cur) {
			n = len(cur)
			pre = string(pre[:len(cur)])
		}

		//循环对比两个字符串
		for j := 0; j < n; j++ {
			if pre[j] == cur[j] {
				continue
			}
			//第一个字符就不相等 则直接返回
			if j == 0 {
				return ""
			}
			pre = string(pre[0:j])
			break
		}
	}
	return pre
}

//字符串查找
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hLen := utf8.RuneCountInString(haystack)
	nLen := utf8.RuneCountInString(needle)
	fmt.Printf("hLen=%d\n", hLen)
	fmt.Printf("nLen=%d\n", nLen)
	//查找字符串的长度大于原字符串 返回-1
	if nLen > hLen {
		return -1
	}

	//是否考虑Unicode字符，如果不考虑可以使用byte
	stack := []rune(haystack)
	need := []rune(needle)
	for i := 0; i < hLen; i++ {
		//不相等则循环下一个字符
		if stack[i] != need[0] {
			continue
		}

		//剩余字符串长度小于查找字符串长度
		if hLen-i < nLen {
			return -1
		}
		if string(stack[i:(nLen+i)]) == needle {
			return i
		}
	}
	return -1
}

//二进制字符串相加
//求字符串的长度utf8.RuneCountInString(a)
//字符类型计算（ASCII值） 字符运算 需要转换0=>48 1=>49
//考虑进位问题
func AddBinary(a, b string) string {
	strAlen := utf8.RuneCountInString(a)
	strBlen := utf8.RuneCountInString(b)
	if strBlen == 0 {
		return ""
	}

	//字符串最大长度
	n := strAlen

	//补全较短的字符串
	if strAlen < strBlen {
		n = strBlen
		//补全A字符串
		a = fmt.Sprintf("%s%s", strings.Repeat("0", strBlen-strAlen), a)
	} else {
		//否则补全B的字符串
		b = fmt.Sprintf("%s%s", strings.Repeat("0", strAlen-strBlen), b)
	}
	res := make([]byte, n+1)

	carry := false
	temp := 0
	for i := n - 1; i >= 0; i-- {

		//字符运算 需要转换0=>48 1=>49
		if carry {
			temp = int(a[i]-'0') + int(b[i]-'0') + 1
		} else {
			temp = int(a[i]-'0') + int(b[i]-'0')
		}

		if temp == 0 {
			res[i+1] = byte('0')
			carry = false
		} else if temp == 1 {
			res[i+1] = byte('1')
			carry = false
		} else if temp == 2 {
			res[i+1] = byte('0')
			carry = true
		} else if temp == 3 {
			res[i+1] = byte('1')
			carry = true
		}
	}

	if carry {
		res[0] = byte('1')
		return string(res)
	} else {
		return string(res[1:])
	}
}

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
