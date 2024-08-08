package sliding_window

import "testing"

/*
找到字符串中所有字母异位词
给定两个字符串s和p找到s中所有p的异位词的子串。
TODO:返回这些子串的起始索引。
不考虑答案输出的顺序.
eg：

	输入: s = "cbaebabacd", p = "abc"
	输出: [0,6]
*/
func findAnagrams(s string, p string) []int {
	/*
			要找到s中p的异位词，首先回忆哈希中的异位词分组，
			会用到哈希表来统计字符出现的次数。然后据此判断是不是异位词。 map[[26]int]string
		解题思路：
			1. 使用两个哈希表分别记录字符串p和滑动窗口中的字符串s的字符出现次数
			2. 使用双指针left和right表示滑动窗口的左右边界
			3. 遍历字符串s，right指针不断右移，每次移动都更新哈希表中的字符出现次数
			4. 当窗口大小等于字符串p的长度时，判断两个哈希表是否相等，如果相等，说明找到了一个异位词
			5. 移动left指针收缩窗口，然后继续遍历字符串s
	*/
	sLen, pLen := len(s), len(p)
	left, right := 0, 0
	// 先初始化两个哈希表，存放每个字符出现的次数
	pMap := [26]int{}
	sMap := [26]int{}
	ret := make([]int, 0)
	// 统计p中每个字符出现的次数
	for _, v := range p {
		pMap[v-'a']++
	}
	// 外层for循环遍历整个字符串s，内层没有循环，是因为内层需要动态维护滑动窗口
	for right < sLen {
		// 扩展右边界
		sMap[s[right]-'a']++
		// 当窗口大小超过p的长度时，收缩左边界
		if right-left+1 > pLen {
			sMap[s[left]-'a']--
			// 移动左指针
			left++
		}
		// 如果窗口大小等于p并且频率与p相同，记录left
		if right-left+1 == pLen && sMap == pMap {
			ret = append(ret, left)
		}
		right++
	}
	return ret
}

func TestName(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	t.Log(result)
}
