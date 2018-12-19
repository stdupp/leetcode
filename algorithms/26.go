func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    store := 1
    pre := nums[0]
    l := 1
    for i:=1; i < len(nums); i++ {
        if nums[i] != pre {
            nums[store] = nums[i]
            store++
            pre = nums[i]
            l++
        }
    }
    return l
}
