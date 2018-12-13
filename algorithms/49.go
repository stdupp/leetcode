import "sort"

func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, str := range strs {
        bs := []byte(str)
        sort.Slice(bs, func (i, j int) bool{return bs[i] < bs[j]})
        s := string(bs)
        groups[s] = append(groups[s], str)
    }
    
    var result [][]string
    for _, v := range groups {
        result = append(result, v)
    }
    return result
}
