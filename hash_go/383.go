package hash

// 救赎金 判断字符串a中的字符是否全部出现在b中（真包含）

func canConstruct(ransomNote string, magazine string) bool {
    var magazineArr [26]int
    for _, c := range magazine {
        magazineArr[c - 'a']++
    }
    for _, c := range ransomNote {
        magazineArr[c - 'a']--
        if magazineArr[c - 'a'] < 0 {
            return false
        }
    }
    return true
}