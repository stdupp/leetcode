import "sort"

func canReorderDoubled(A []int) bool {
    m := make(map[int]int, len(A))
    for i, v := range A {
        if v < 0 {
            v = -v
            A[i] = v
        }
        m[v]+=1
    }
    sort.Ints(A)
    for _, v := range A {
        if c :=  m[v]; c == 0 {
            continue
        }
        if c := m[2*v]; c == 0 {
            return false
        }
        m[v]--
        m[2*v]--
    }
    return true
}
