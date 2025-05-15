// func reverseVowels(s string) string {
//     runes := []rune(s)
//     vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
//     vowelsInS := []rune{}
//     for _, c := range s {
//         if slices.Contains(vowels, c) {
//             vowelsInS = append(vowelsInS, c)
//         }
//     }
//     j := len(vowelsInS)-1
//     for i, c := range s {
//         if slices.Contains(vowels, c) {
//             runes[i] = vowelsInS[j]
//             j--
//         }
//     }
//     return string(runes)
// }

func reverseVowels(s string) string {
    i := 0
    j := len(s) - 1
    runes := []rune(s)
    vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
    for i < j {
        if slices.Contains(vowels, runes[i]) {
            for !slices.Contains(vowels, runes[j]) {
                j--
            }
            tmp := runes[i]
            runes[i] = runes[j]
            runes[j] = tmp
            j--
        }
        i++
    }
    return string(runes)
}