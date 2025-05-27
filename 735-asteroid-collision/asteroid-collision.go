func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, v := range asteroids {
		stack = collate(stack, v)
	}
	return stack
}

func isCollision(last int, v int) bool {
	return last > 0 && v < 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func collate(stack []int, v int) []int {
	if len(stack) == 0 {
		return append(stack, v)
	}

	last := stack[len(stack)-1]

	if !isCollision(last, v) {
		return append(stack, v)
	}

	if abs(last) == abs(v) {
		return stack[:len(stack)-1]
	} else if abs(last) > abs(v) {
		return stack
	} else {
		return collate(stack[:len(stack)-1], v)
	}
}
