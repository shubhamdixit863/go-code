package main

import (
	"fmt"
	// "go-code/directions"
	// "go-code/fun"
	// "go-code/fizzbuzz"

)

// func main() {
// 	fmt.Println("Directions:")
// 	fmt.Println("North:", directions.North)
// 	fmt.Println("South:", directions.South)
// 	fmt.Println("East:", directions.East)
// 	fmt.Println("West:", directions.West)
// }


// func main() {
// 	sum, diff, product := fun.Calculate(10, 2, 3, 4, 5)

// 	fmt.Println("Sum:", sum)
// 	fmt.Println("Difference:", diff)
// 	fmt.Println("Product:", product)
// }
// func main(){
// 	fmt.Println(fizzbuzz.Fizzbuzz(3))
// 	fmt.Println(fizzbuzz.Fizzbuzz(5))
// 	fmt.Println(fizzbuzz.Fizzbuzz(10))
// 	fmt.Println(fizzbuzz.Fizzbuzz(15))

// }
// func main() {

// 	fizzBuzz := func(num int) {
// 		if num%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if num%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(num)
// 		}
// 	}

// 	for i := 1; i <= 10; i++ {
// 		fizzBuzz(i)
// 	}
// }
func applyOperation(operation func(int) int) {

	// loop from 1 to 5
	for i := 1; i <= 5; i++ {
		result := operation(i)
		fmt.Println(result)
	}
}

// func main() {

// 	// passing anonymous function
// 	applyOperation(func(n int) int {
// 		if n%2 == 0 {
// 			return n * 2
// 		} else {
// 			return n * 3
// 		}
// 	})
// }