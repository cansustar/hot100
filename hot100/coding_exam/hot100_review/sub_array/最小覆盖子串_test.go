package subarray

import "testing"

/*
最小覆盖子串
给你一个字符串 s 、一个字符串 t 。
返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

eg:
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
*/

func MinCoverSubarray(strs string, t string) string {
	/*
		slide_windows
		定义两个哈希表，一个是need，一个是window
		need存储目标字符串t中，每个字符出现的次数
		window存储每个窗口中，每个字符出现的次数
	*/
	need := make(map[byte]int)
	window := make(map[byte]int)
	minSub := ""
	// 初始化need
	for _, v := range t {
		need[byte(v)]++
	}
	// 定义左右指针
	left, right := 0, 0
	// 定义窗口中满足need条件的字符个数
	valid := 0
	for right < len(strs) {
		// 扩展右边界
		c := strs[right]
		right++
		// 更新窗口中的数据
		if need[c] > 0 {
			window[c]++
			// 如果窗口中的字符个数等于need中的字符个数，说明已经有一个字符满足条件
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否需要收缩。 当need中的字符个数等于valid时，说明窗口中已经包含了所有的字符
		for valid == len(need) {
			// 更新最小覆盖子串 当当前窗口大小小于最小子串长度时，或者最小子串为空时，更新最小子串
			if right-left < len(minSub) || minSub == "" {
				minSub = strs[left:right]
			}
			// 收缩左侧窗口 d为待删除的字符
			d := strs[left]
			left++
			// 更新窗口数据。 如果d在need中，且窗口中的字符个数等于need中的字符个数，valid减1
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return minSub
}

func TestMinSubarray(t *testing.T) {
	strs := "ADOBECODEBANC"
	tStr := "ABC"
	result := MinCoverSubarray(strs, tStr)
	t.Log(result)
}
