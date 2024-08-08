package subarray

import (
	"fmt"
)

// buildPrefixSum 构建前缀和数组
func buildPrefixSum(arr []int) []int {
	n := len(arr)
	// 为了避免数组越界，prefixSum 长度为 n+1
	// prefixSum[i] 表示 arr[0] 到 arr[i-1] 的和
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// prefixSum[i-1]=arr[0]+...+arr[i-2],再加上arr[i-1]就是prefixSum[i]
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
	}
	// 构建好的前缀表
	return prefixSum
}

// rangeSum 计算区间 [left, right] 的和
func rangeSum(prefixSum []int, left, right int) int {
	return prefixSum[right+1] - prefixSum[left]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	prefixSum := buildPrefixSum(arr)
	fmt.Println(prefixSum) // 输出: [0 1 3 6 10 15]

	left, right := 1, 3
	fmt.Println(rangeSum(prefixSum, left, right)) // 输出: 9 (2 + 3 + 4)
}
