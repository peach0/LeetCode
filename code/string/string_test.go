package stringCode

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	str := ""
	t.Log(lengthOfLongestSubstring(str))
}
//3. 重复字符的最长子串
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
func TestPalindRome(t *testing.T) {
	s := "aaaa"
	print(longestPalindrome(s))
}
//5. 最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	start := 0
	max   := 1
	dp := make([]bool, n)

	for r := 0; r < n; r++ {
		dp[r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l  <= 2 || dp[l+1]) {
				dp[l] = true
				if r - l + 1 > max {
					start = l
					max   = r - l + 1
				}
			}else {
				dp[l] = false
			}
		}
	}
	return  s[start:start+max]
}

func TestConvert(t *testing.T) {
	s := "ABCD"
	println(convert(s, 3))
}
//6. Z 字形变换
//LCIRETOESIIGEDHN
//LCIRETOESIIEDHN
func convert(s string, numRows int) string {
	if s == "" || len(s) <=  numRows || numRows <= 1 {
		return s
	}
	var b []byte
	for i := 0; i < numRows ; i ++ {
		j := i
		for j< len(s) {
			b = append(b, s[j])
			if i != 0 && i != numRows -1 && 2 * numRows + j - 2 * i - 2 < len(s){
				b = append(b, s[2 * numRows + j - 2 * i - 2])
			}
			j = j + 2 * numRows - 2
		}
	}
	return string(b)
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}