func rotate(nums []int, k int)  {
    if len(nums) == 0{
        return
    }
    
    step := k %  len(nums)
    start_index := len(nums) - step
    
    tmp := append(nums[start_index:], nums[:start_index]...)
    for i, _:= range nums{
        nums[i] = tmp[i]
    }
}
