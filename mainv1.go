package main

import "fmt"

// porvided by go library ,you dont need to invoke it by yourslef ,it gets called automatically before main
func init() {
	fmt.Println("I will be called before main")
}

func main() {
	fmt.Println("Main is called")

	//// annonymous function or lmabda function they dnt need to have name
	//c := func(param int) int {
	//	/// you can put logic
	//	fmt.Println(param)
	//	return 2 * 2
	//}(2)

	//fmt.Println(c)

	//func() {
	//
	//}()

	//c := foo()
	//
	//// c is also a function
	//
	//c()

	applyOperation(func(x int) int {
		if x%2 == 0 {
			return x * 2
		} else {
			return x * 3
		}

	})

	gh(2)
}

// function foo is a higher order  function because it is returning a function
func foo() func() int {

	// i can return a function from it
	return func() int {
		return 2
	}
}

// bar is a higher order function as well
// param is a call back function
func bar(param func(p int) int) {
	param(2)
}

/**
Write a Go program that demonstrates a higher-order function.
 1. Create a higher-order function named applyOperation that:
 • Accepts another function as a parameter
 • The passed function should take an integer and return an integer
 2. Inside applyOperation:
 • Use a for loop to iterate from 1 to 5
 • For each number, call the passed function
 • Print the returned result
 3. In main:
 • Pass an anonymous function to applyOperation
 • The anonymous function should use if–else logic:
 • If the number is even, return n * 2
 • If the number is odd, return n * 3

*/

func gh(p int) {

}

func hj(k func(ji int) int) {

}
func applyOperation(p func(x int) int) {
	for i := 1; i <= 5; i++ {
		ret := p(i)
		fmt.Println(ret)
	}
}
