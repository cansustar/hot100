package double_pointer

import "testing"

/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
eg:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

func MoveZero(nums []int) []int {
	// 双指针法
	// 一个指针用于遍历数组，另一个指针用于指向非0元素
	// 遍历数组，如果当前元素不为0，则将其与非0指针指向的元素交换
	// 时间复杂度O(n), 空间复杂度O(1)
	// 初始化非0指针
	j := 0
	// 指针i用于遍历数组，指针j用于指向非0元素
	for i := 0; i < len(nums); i++ {
		// 如果i指向元素不为0，则与非0指针指向的元素交换
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			// i指向元素非0.同时移动非0指针j
			j++
		}
		// 当i指向元素为0的时候，不移动j. 这样就保证了非0元素的相对顺序
	}
	return nums
}

func MoveZero2(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	result := MoveZero(nums)
	t.Log(result)

}
