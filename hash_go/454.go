package hash

// 四数之和
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    numMap := make(map[int]int)
    aggregate := 0
    for _, i := range nums1 {
        for _, j := range nums2 {
            numMap[i+j]++
        }
    }
    for _, i := range nums3 {
        for _, j := range nums4 {
            if n, ok := numMap[-i-j]; ok {
                aggregate += n
            }
        }
    }
    return aggregate
}