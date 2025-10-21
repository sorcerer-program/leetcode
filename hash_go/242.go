package hash

// 判断是否有效字母异位词

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var hash [26]int
    for _, c := range s {
        hash[c - 'a']++
    }
    for _, c := range t {
        hash[c - 'a']--
    }
    for _, c := range hash {
        if c != 0 {
            return false
        }
    }
    return true
}