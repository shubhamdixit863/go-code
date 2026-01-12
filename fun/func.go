package fun
func Calculate(a, b, c, d, e int) (int, int, int) {
	sum := a + b + c + d + e
	diff := a - b - c - d - e
	product := a * b * c * d * e

	return sum, diff, product
}