func removeElement(nums []int, val int) int {    
    pre := -1
    var length int
    for i, v := range nums {
        if v != val {
            length++
            pre++
            if pre != i {
                nums[pre] = nums[i]
            }
        }
    }
    return length
}
