package highorder
import "fmt"

func ApplyOperation(operation func(int) int) {
	for i := 1; i <= 5; i++ {
		result := operation(i)
		fmt.Println("Input:", i, "Output:", result)
	}
}