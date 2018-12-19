func getHint(secret string, guess string) string {
    var bull int
    var total int
    
    var list1  [10]int
    for _, v := range secret {
        i := int(v - '0')
        list1[i]++
    }
    
    for i, v := range guess {
        if secret[i] == guess[i] {
            bull++
        }
        if list1[int(v - '0')] > 0 {
            total++
            list1[int(v - '0')]--
        }
        
    }
    
    return fmt.Sprintf("%dA%dB", bull, total - bull)
    
    
}
