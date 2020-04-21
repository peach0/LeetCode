package stringCode

import (
	"net/http"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	str := ""
	t.Log(lengthOfLongestSubstring(str))
}
//3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	start, len := 0, 0
	for end := 0; end < n; end++ {
		if v, ok := m[s[end]]; ok {
			start = Max(start, v)
		}
		len = Max(len, end - start + 1)
		m[s[end]] = end + 1
	}
	return  len
}
func TestStr(t *testing.T) {
	str := "hahahah"
	t.Log(str[1:])

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//5. 最长回文子串
func longestPalindrome(s string) string {
}
