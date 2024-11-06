func isBalanced(num string) bool {
    evenSum := 0
    oddSum := 0
    for i, c := range num {
        n := int(c - '0')
        if i % 2 == 0 {
            evenSum += n
        } else {
            oddSum += n
        }
    }
    return evenSum == oddSum
}