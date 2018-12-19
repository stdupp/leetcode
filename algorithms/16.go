import "sort"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    var mini_diff int
    var sum int
    for i := 0; i < len(nums) -2; i++ {
        l, r := i+1, len(nums) -1
        for l < r {
            tmp := nums[i] + nums[l] + nums[r]
            diff := tmp -target 
            if diff == 0 {
                    return tmp
            } else if diff > 0 {
                    if diff < mini_diff || mini_diff == 0 {
                            mini_diff = diff
                            sum = tmp
                    }
                    r--
            } else {
                    diff = -diff
                    if diff < mini_diff || mini_diff == 0 {
                            mini_diff = diff
                            sum = tmp
                    }
                    l++
            }
        }
    }
    
    return sum
}
