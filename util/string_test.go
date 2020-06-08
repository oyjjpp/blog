package util

import "testing"

func TestVersionCompare(t *testing.T) {
	res := VersionCompare("52.1", "51.1", ">=")
	t.Log(res)
	res1 := VersionCompare("102", "51.1", ">")
	t.Log(res1)
	res2 := VersionCompare("52.1", "51.1", "<")
	t.Log(res2)
	res3 := VersionCompare("102", "51.1", "<=")
	t.Log(res3)
	res4 := VersionCompare("51.1.0", "51.1", "==")
	t.Log(res4)
}
