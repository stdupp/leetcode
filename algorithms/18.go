import "sort"

func fourSum(nums []int, target int) [][]int {
    var result [][]int
    sort.Ints(nums)
    l := len(nums)
    
    for i :=0 ; i < l -3; i++{
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        
        if 4 * nums[i] > target {
            break
        }
        
        for j := l-1;  j > i+2; j-- {
            if j != l-1 && nums[j] == nums[j+1] {
                continue
            }
            
            if 4 * nums[j] < target {
                break
            }
            
            sub := target - nums[i] - nums[j]
            low, high := i+1, j-1
            for low < high {
                sum := nums[low] + nums[high]
                if sum > sub {
                    high--
                } else if sum < sub {
                    low++
                } else {
                    result = append(result, []int{nums[i], nums[low], nums[high], nums[j]})
                    for low < high && nums[high] == nums[high -1] {
                        high--
                    }
                    
                    for low < high && nums[low] == nums[low+1] {
                        low++
                    }
                    low, high = low+1, high -1
                }
            }
        }
    }
    
    return result
}
