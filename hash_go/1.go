package hash
// 两数之和
func twoSum(nums []int, target int) []int {
  set := make(map[int]int)
  for index, val := range nums {
    if i, ok := set[target - val]; ok {
      return []int{index, i}
    }else {
      set[val] = index
    }
  }
  return []int{}
}