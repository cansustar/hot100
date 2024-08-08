package 哈希

import (
	"fmt"
	"testing"
)

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

func TwoSum(nums []int, target int) []int {
	// 哈希表的方法
	// key: 数组元素的值, value: 数组元素的下标
	hashtable := make(map[int]int)
	for i, v := range nums {
		// 如果 target-v 在哈希表中存在, 则返回
		if j, ok := hashtable[target-v]; ok {
			return []int{i, j}
		} else {
			hashtable[v] = i
		}
	}
	return nil
}

// 测试
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Println(result)
}
