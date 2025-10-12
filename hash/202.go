package hash
// 判断是否是快乐数
func isHappy(n int) bool {
  set := make(map[]struct{})
  for n != 1 {
    n = getSum(n)
    if _, ok = set[n]; ok {
      return false
    }
  }
  return true
}
func getSum(n int) {
  sum := 0
  for n > 0 {
    remainder = n % 10
    sum + = remainder * remainder
    n /= n
  }
  return sum
}