func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    ms := make(map[byte]byte)
    mt := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        valt, existt := ms[s[i]]
        vals, exists := mt[t[i]]
        if !existt {
            ms[s[i]] = t[i]
        }
        if !exists {
            mt[t[i]] = s[i]
        }
        if (existt && valt != t[i]) || (exists && vals != s[i]) {
            return false
        }
    }
    return true
}