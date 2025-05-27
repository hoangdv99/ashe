func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i, v := range asteroids {
		if i == 0 {
			stack = append(stack, v)
			continue
		}
		var last int
        if len(stack) > 0 {
            last = stack[len(stack)-1]
        }
		if !isCollision(last, v) {
			stack = append(stack, v)
		} else {
			stack = collate(stack, v)
		}
	}
	return stack
}

func isCollision(last int, v int) bool {
	return last > 0 && v < 0
}

func collate(stack []int, v int) []int {
	var last int
	if len(stack) > 0 {
		last = stack[len(stack)-1]
	}
	if !isCollision(last, v) {
		return append(stack, v)
	}
	if math.Abs(float64(last)) == math.Abs(float64(v)) {
		return stack[:len(stack)-1]
	} else if math.Abs(float64(last)) > math.Abs(float64(v)) {
		return stack
	} else {
		return collate(stack[:len(stack)-1], v)
	}
}
