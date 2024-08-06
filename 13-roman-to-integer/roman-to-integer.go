func romanToInt(s string) int {
    symbols := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var total int
	i := len(s) - 1

	for i >= 0 {
		if i > 0 {
			value, exist := symbols[string(s[i-1])+string(s[i])]
			if exist {
				total += value
				i -= 2
				continue
			}
		}
		total += symbols[string(s[i])]
		i--
	}
	return total  
}