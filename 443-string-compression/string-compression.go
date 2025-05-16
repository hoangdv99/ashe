func compress(chars []byte) int {
	read := 0
    write := 0
    for read < len(chars) {
        currentChar := chars[read]
        count := 0
        for read < len(chars) && chars[read] == currentChar {
            count++
            read++
        }
        chars[write] = currentChar
        write++
        if count > 1 {
            for _, digit := range strconv.Itoa(count) {
                chars[write] = byte(digit)
                write++
            }
        }
    }
    return write
}
