func majorityElement(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    
    num1, num2 := nums[0], nums[0]
    count1, count2 := 0, 0
    
    for _, num := range nums {
        if num == num1 {
            count1++
            continue
        }
        
        if num == num2 {
            count2++
            continue
        }
        
        if count1 == 0 {
            num1 = num
            count1++
            continue
        }
        
        if count2 == 0 {
            num2 = num
            count2++
            continue
        }
        
        count1--
        count2--
    }
    
    count1, count2 = 0, 0
    
    for _, num := range nums {
        if num == num1 {
            count1++
        } else if num == num2 {
            count2++
        }
    }
    var result []int
    if count1 > len(nums)/3 {
        result = append(result, num1)
    }
    
    if count2 > len(nums)/3 {
        result = append(result, num2)
    }
    
    return result
}
