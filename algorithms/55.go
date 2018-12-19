func canJump(nums []int) bool {
    index := len(nums) -1
    tmp := index
    for index > 0 && tmp >0 {
        tmp -= 1
        if nums[tmp] >= (index - tmp) {
            index = tmp
        }
        
    }
    
    if index <= 0 {
        return true
    }
    return false
}
