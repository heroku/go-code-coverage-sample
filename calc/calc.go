package calc

func Add(a ...int) int {
	var sum int
	for _, a := range a {
		sum += a
	}
	return sum
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a ...int) int {
	prod := 1
	for _, a := range a {
		prod *= a
	}
	return prod
}

func Divide(a, b int) int {
	return a / b
}
