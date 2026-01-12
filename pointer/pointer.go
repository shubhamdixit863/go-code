package pointer

// This function takes two integer pointers
func SwapAndAdd(a *int, b *int) {

	temp := *a
	*a = *b
	*b = temp

	*a = *a + 10
	*b = *b + 10
}
