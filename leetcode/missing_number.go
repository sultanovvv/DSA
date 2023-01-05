package leetcode

func missingNumber(nums []int) int {
	var sumF, sum int

	for i := 1; i <= len(nums); i++ {
		sum += i
		sumF += nums[i-1]
	}
	return sum - sumF
}
