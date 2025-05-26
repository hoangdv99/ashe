func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	freq1 := make(map[byte]int)
	freq2 := make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]]++
		freq2[word2[i]]++
	}

	freqArr1 := []int{}
	freqArr2 := []int{}

	for c := range freq1 {
		freqArr1 = append(freqArr1, freq1[c])
		freqArr2 = append(freqArr2, freq2[c])
	}

	sort.Ints(freqArr1)
	sort.Ints(freqArr2)
	if len(freqArr1) != len(freqArr2) {
		return false
	}
	for i := range freqArr1 {
		if freqArr1[i] != freqArr2[i] {
			return false
		}
	}
	return true
}
