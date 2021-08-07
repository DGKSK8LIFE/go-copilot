// Write a function that returns the sum of the two numbers.
func sum(a, b int) int {
	return a + b
}

// Write a function that returns the product of the two numbers.
func product(a, b int) int {
	return a * b
}

// Write a function that returns the square of the number.
func square(a int) int {
	return a * a
}

// Write a function that multiples two two-dimensional arrays together and returns the result.
func multiply(a [][]int, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := range result {
		result[i] = make([]int, len(a[i]))
		for j := range result[i] {
			result[i][j] = a[i][j] * b[i][j]
		}
	}
	return result
}

// Write a function that returns the maximum of two numbers.
func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Write a function that returns the minimum of two numbers.
func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Write a function that returns the factorial of the number.
func factorial(a int) int {
	if a < 0 {
		panic("factorial of negative number")
	}
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

// Write a function that returns the sine of the number.
func sine(a int) int {
	if a < -90 || a > 90 {
		panic("sine of out of bounds number")
	}
	return int(math.Sin(float64(a) * math.Pi / 180))
}

// Write a function that returns the cosine of the number.
func cosine(a int) int {
	if a < -90 || a > 90 {
		panic("cosine of out of bounds number")
	}
	return int(math.Cos(float64(a) * math.Pi / 180))
}

// Write a function that returns the tangent of the number.
func tangent(a int) int {
	if a < -90 || a > 90 {
		panic("tangent of out of bounds number")
	}
	return int(math.Tan(float64(a) * math.Pi / 180))
}
