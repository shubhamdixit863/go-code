package errorHandling
import "errors"

func Divide(A, B float64) (float64, error) {
	if B == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return A / B, nil
}

