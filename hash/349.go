package hash
// 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
  nums1Set := make(map[int]struct{})
  answerSet := make(map[int]struct{})
  for _, val := range nums1 {
    nums1Set[val] = struct{}{}
  }
  for _, val := range nums2 {
    if _, ok := nums1Set[val]; ok {
      answerSet[val] = struct{}{}
    }
  }
  answerSlice := make([]int, 0, len(answerSet))
  for k := range answerSet {
    answerSlice = append(answerSlice, k)
  }
  return answerSlice
}