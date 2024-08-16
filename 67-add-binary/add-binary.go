func addBinary(a string, b string) string {
	var sum string
	carry := "0"

	if len(a) > len(b) {
		return addBinary(b, a)
	} else if len(a) < len(b) {
		dif := len(b) - len(a)
		for i := 0; i < dif; i++ {
			fmt.Println(i)
			a = "0" + a
		}
	}

	fmt.Println(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		if string(a[i]) == "0" && string(b[i]) == "0" {
			if carry == "0" {
				sum = "0" + sum
			} else {
				sum = "1" + sum
				carry = "0"
			}
		} else if (string(a[i]) == "0" && string(b[i]) == "1") || (string(a[i]) == "1" && string(b[i]) == "0") {
			if carry == "0" {
				sum = "1" + sum
			} else {
				sum = "0" + sum
				carry = "1"
			}
		} else {
			if carry == "0" {
				sum = "0" + sum
				carry = "1"
			} else {
				sum = "1" + sum
				carry = "1"
			}
		}
	}

	if carry == "1" {
		sum = "1" + sum
	}

	return sum
}