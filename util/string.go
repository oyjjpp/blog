package util

import "strings"

// VersionCompare
// 版本比较
func VersionCompare(v1, v2, operator string) bool {
	com := compare(v1, v2)
	switch operator {
	case "==":
		if com == 0 {
			return true
		}
	case "<":
		if com == -1 {
			return true
		}
	case ">":
		if com == 1 {
			return true
		}
	case "<=":
		if com == 0 || com == -1 {
			return true
		}
	case ">=":
		if com == 0 || com == 1 {
			return true
		}
	}
	return false
}

// compare
// 1 : version1 > version2
// -1 : version1 < version2
// 0 : 其他情况
func compare(version1, version2 string) int {
	// 替换一些常见的版本符号
	replaceMap := map[string]string{"V": "", "v": "", "-": "."}
	//keywords := {"alpha,beta,rc,p"}
	for k, v := range replaceMap {
		if strings.Contains(version1, k) {
			strings.Replace(version1, k, v, -1)
		}
		if strings.Contains(version2, k) {
			strings.Replace(version2, k, v, -1)
		}
	}

	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	l := len(v1)
	for i := 0; i < l; i++ {
		vs1 := strings.TrimLeft(v1[i], "0")
		vs2 := strings.TrimLeft(v2[i], "0")

		for len(vs1) < len(vs2) {
			vs1 = "0" + vs1
		}
		for len(vs2) < len(vs1) {
			vs2 = "0" + vs2
		}

		vl := len(vs1)
		for j := 0; j < vl; j++ {
			if vs1[j] == vs2[j] {
				continue
			}

			if vs1[j] < vs2[j] {
				return -1
			}

			return 1
		}
	}

	return 0
}
