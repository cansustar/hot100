package hash

import (
	"fmt"
	"testing"
)

// 异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。返回分组后的结果
//字母异位词指字母相同，但排列不同的字符串。

func groupAnagrams(strs []string) [][]string {
	// 统计每个字符串中，各个字符出现的次数，将其存到哈希表中。
	// key: 字符串中各个字符出现的次数, value: 字符串
	// 最后返回哈希表中的所有 value 即可
	// 时间复杂度O(n*k), n为字符串数组长度, k为字符串的平均长度
	// 空间复杂度O(n)

	// 用26个长度的数组作为key,表示各个字母出现的次数。 value为字符串的切片
	hashTable := make(map[[26]int][]string)
	for _, str := range strs {
		// cnt用来记录该字符串，每个字符出现的次数
		cnt := [26]int{}
		// 遍历当前字符串的每个字符，更新cnt列表
		for _, c := range str {
			cnt[c-'a']++
		}
		// // 将str加入到出现字符数量一样的字符串切片中
		hashTable[cnt] = append(hashTable[cnt], str)
	}
	ans := make([][]string, 0, len(hashTable))
	for _, v := range hashTable {
		ans = append(ans, v)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
