func mergeAlternately(word1 string, word2 string) string {
    var n int
    if len(word1) >= len(word2) {
        n = len(word1)
    } else {
        n = len(word2)
    }

    var result []byte
    for i := 0; i < n; i++ {
        if i < len(word1) {
            result = append(result, word1[i])
        }
        if i < len(word2) {
            result = append(result, word2[i])
        }
    }

    return string(result)
}