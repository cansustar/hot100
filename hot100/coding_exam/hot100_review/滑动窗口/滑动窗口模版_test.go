package sliding_window

func slidingWindow(nums []int, k int) []int {
	// n是数组的长度， k是窗口大小
	n := len(nums)
	// 如果n或者k为0，这种情况下没有可处理的数据或窗口，直接返回空数组
	if n*k == 0 {
		return []int{}
	}
	// 如果窗口大小为1，说明窗口每次只覆盖一个元素，这样与遍历整个数组的效果一样，所以直接返回原数组
	if k == 1 {
		return nums
	}
	// 滑动窗口实际上就是某个数据结构+双指针
	left, right := 0, 0
	// 数据结构，用于存储窗口中的元素和相关信息
	window := make(map[int]int)
	var result []int

	for right < n {
		// 窗口右边界扩展。 TODO 这里可以添加一些扩展逻辑
		window[nums[right]]++
		right++

		// 当窗口大小等于k时，(或者遇到了其他需要暂停拓展的条件)，处理窗口内的元素
		if right-left == k {
			// 处理窗口中的元素，例如计算最大值、最小值、或特定条件下的子数组
			result = append(result, processWindow(window))

			// 窗口左边界收缩  收缩不仅意味着指针的移动，还需要对数据结构进行相应的更新。
			// 如果窗口左边界的元素出现次数为0，说明窗口中不再有该元素，可以删除
			delete(window, nums[left])
			// 移动指针，即窗口左边界右移
			left++
		}
	}

	return result
}
