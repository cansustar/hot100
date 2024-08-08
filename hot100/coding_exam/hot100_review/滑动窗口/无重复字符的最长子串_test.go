package sliding_window

import (
	"math"
	"testing"
)

/*
无重复字符的最长子串
给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。
eg:
	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func lengthOfLongestSubString(s string) int {
	/*
		要统计不重复的子串的最大长度，一个是最大长度，一个是不重复。
		做到不重复，可以用哈希表来判断是否重复。
		做到找到字串的最大长度，可以用滑动窗口。
			其实就是指针，一个指向子串的起始位置，一个指向子串的结束位置。
			判断右指针指向的字符是否在子串中，如果在，左指针右移，直到该元素不在子串中。再移动右指针。
	*/
	n := len(s)
	left, right := 0, 0
	// 窗口，记录字符串出现过
	windows := make(map[byte]bool)
	res := 0
	// 右指针小于字符串长度 这里是两个for循环，可以理解为，外层for循环遍历整个字符串，内层for循环处理窗口
	for right < n {
		// 扩展右边界条件：右指针小于字符串长度，且窗口中不包含当前字符
		for right < n && !windows[s[right]] {
			windows[s[right]] = true
			right++
		}
		// 窗口右侧出现重复元素，停止扩展，处理当前最大长度
		cur_len := right - left
		res = int(math.Max(float64(res), float64(cur_len)))
		// 当窗口中出现重复元素时，收缩窗口
		delete(windows, s[left])
		left++

	}
	return res
}

func TestLengthOfLongestSubString(t *testing.T) {
	s := "abcabcbb"
	result := lengthOfLongestSubString(s)
	t.Log(result)
}
