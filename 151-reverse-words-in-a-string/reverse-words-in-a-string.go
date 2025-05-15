func reverseWords(s string) string {
    words := strings.Split(s, " ")
    reverseWordList := []string{}
    for i := len(words)-1; i >= 0; i-- {
        if words[i] != "" {
            reverseWordList = append(reverseWordList, words[i])
        }
    }
    return strings.Join(reverseWordList, " ")
}