package leetcode

func removeDuplicates(nums []int) int {

	tmp := nums[0]
	for i := 1; i < len(nums); {
		if tmp == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		} else if tmp < nums[i] {
			tmp = nums[i]
			i++
		}
	}
	return len(nums)
}
