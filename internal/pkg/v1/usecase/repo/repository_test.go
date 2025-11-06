package repo

import (
	"regexp"
	"strings"
	"testing"

	"github.com/mozillazg/go-pinyin"
)

func TestCreateAnchorUser(t *testing.T) {
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	result := pinyin.Pinyin("国内-瑞雪验收-112", a)
	firstLetters := make([]string, len(result))
	for i, item := range result {
		firstLetters[i] = strings.Join(item, "")
	}
	name := strings.Join(firstLetters, "")
	t.Log(clearString(name))

}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
