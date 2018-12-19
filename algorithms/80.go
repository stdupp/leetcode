func removeDuplicates(nums []int) int {
    pre := -1
    length := 0
    var tmp int
    var num int
    for i, v := range nums {
        if i == 0 {
            tmp = v
            num = 1
            length++
            pre++
            continue
        }
        
        if v == tmp {
            if num < 2 {
                length++
                num++
                pre++
                if pre != i {
                    nums[pre] = nums[i]
                }   
            } else {
                num++
            }
        } else {
            pre++
            length++
            if pre != i {
                nums[pre] = nums[i]
            }
            tmp = v
            num = 1
        } 
        
    }
    return length
}
