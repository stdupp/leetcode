import "strings"

func wordPattern(pattern string, str string) bool {

    slist := strings.Split(str, " ")
    if len(slist) != len(pattern) {
        return false
    }
    
    words := make(map[byte]string)
    ps := make(map[string]byte)
    
    for i, tmp :=range pattern {
        b := byte(tmp)
        w, ok := words[b]
        if !ok {
            words[b] = slist[i]
        } else {
            if slist[i] != w {
                return false
            }
        }
        
        p, ok := ps[slist[i]]
        if !ok {
            ps[slist[i]] = b
        } else {
            if p != b {
                return false
            }
        }
    }
    
    return true
}
