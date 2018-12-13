var result [][]int

func subsets(nums []int) [][]int {
    result = make([][]int, 0, 3)
    findSet([]int{}, nums)
    return result
}

func findSet(set []int, nums []int) {
    result = append(result, set)
    
    for i, num := range nums {
        tmp := make([]int, len(set)+1)
        copy(tmp, set)
        tmp[len(set)] = num
        findSet(tmp, nums[i+1:])
    }
}
