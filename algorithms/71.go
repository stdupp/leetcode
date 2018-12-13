import "strings"

func simplifyPath(path string) string {
    paths := strings.Split(path, "/")
    dirs := []string{}
    for _, p := range paths {
        if p == ".." {
            if len(dirs) > 0 {
                dirs = dirs[:len(dirs) - 1]
            }
        } else if p == "." || p == "" {
            continue
        } else {
            dirs = append(dirs, p)
        }
    }
    
    return "/" + strings.Join(dirs, "/")
}
