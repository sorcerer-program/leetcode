package hash

// 三数之和，这道题不适合用hash

func threeSum(nums []int) [][]int {
	var answer [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return answer
		}
		for nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				for nums[right] == nums[right-1] {
					right--
				}
				right--
			}else if sum < 0 {
				for nums[left] == nums[left+1] {
					left++
				}
				left++
			}else {
				answer = append(answer, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return answer
}