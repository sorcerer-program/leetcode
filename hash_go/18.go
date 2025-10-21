package hash

// 四数之和
func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	answer := [][]int{}
	if size < 4 {
		return answer
	}
	for i := 1; i < size-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < size-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, size-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					for nums[left] == nums[left+1] {
						left++
					}
					for nums[right] == nums[right-1] {
						right--
					}
					right++
					left--
				}
			}
		}
	}
	return answer
}