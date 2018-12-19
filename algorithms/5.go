func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    
    max := string(s[0])
    for i := 0; i < len(s); i++ {
        max = calPa(s, i, i, max)
        max = calPa(s, i, i+1, max)
    }
    
    return max
}

func calPa(s string, i, j int, max string) string {
    var sub string

    for i >=0 && j <= len(s) -1 && s[i] == s[j] {
        sub = s[i:j+1]
        i--
        j++
    }
    
    if len(max) < len(sub) {
        max = sub
    }
    
    return max
}
