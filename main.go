package main

import (
	"fmt"
	"os"
)

// g:=9  // cannot do this outside the function
var g = 9 // global variable

// define multiple variable definitions
var (
	a int    = 3
	x string = "Hello world"
)

const alpha = 90 // once they are declared you cannot change them

func main() {

	// Print all arguments, including the program path
	fmt.Println("All arguments:", os.Args)

	// Print only user-provided arguments
	argsWithoutProg := os.Args[1:]
	fmt.Println("User arguments:", argsWithoutProg)

	// Access a specific argument
	if len(argsWithoutProg) > 0 {
		fmt.Println("First user argument:", argsWithoutProg[0])
		fmt.Println("Second user argument:", argsWithoutProg[1])
	}

	// add method

	// subtract methd
	//var c bool
	//var c string // uniocde (utf-8),asciii

	/**
	int  int8  int16  int32  int64  // it can store both positivce and negative
	uint uint8 uint16 uint32 uint64 uintptr  // it can only store insiogned values that is positive values
	*/

	//var z int32

	//var z uint8 // it can store an integer that can be respreented in 8 bits
	//z = -127    // the maximum value that it cna store is from -127 to 127
	//fmt.Println(z)
	// uintptr is for storing pointer variables
	//var d uint8
	//d = 9
	//var c uintptr
	//c = d

	// byte // 1 byte =8 bits
	//
	//var c byte  // []byte slice
	//
	//c = 9
	// rune  // alias for int32
	//var c rune
	//c = 90

	//float32 float64

	//var t float32
	//t = 8999999999999999999999999999999999999999.900000000000000000000000000

	// complex64 complex128 // 2+3i (-1)
	//var z1 complex128 = 3 + 4i
	//var z int8 = 2 // [10] ---> // [00000010]
	//fmt.Println(z)

	//var s float32
	////fmt.Println(s)
	////var z =9
	////var x =&z  // x is basically a pointer  // default valu of pointer is nil
	//
	////var t *int
	////fmt.Println(t)
	//s = 90
	//fmt.Println(s)

	//var c int
	//c = 9 // assignment opertaor
	// shorthand notation
	//c:=9 // compiler is smart enough to understand that the data type of c is 9
	//var g = 9 // the compiler is smart enough to understand the type

	var e = 5.5
	var d = false

	fmt.Println(d, e)
	a = 90

	//f := Sum(9, 8) // arguments
	//fmt.Println(f)

	//conditionals(3)
	//cond(5)

	//switchCases(20)

	//Loop()
	//	evenOdd()
}

// You have to setup a go app and with   one folder and it  should be called as
// directions it will have a file called called direction.go and it will contain all the const values
// of directions for ex const North=1 and so on
// print these directions in main.go file

// void ff
// A function is a set of encapsulated code that you can re use as many times as you want

func Sum(a int, b int) (int, int, int) { // declaration and definition
	// is where logic goes
	sum := a + b
	//fmt.Println(sum)
	return sum, a - b, a * b
}

// parameter and arguments

// function call // where you are actually calling the function

// Now create a function that takes 5 int parameters   in the package called as fun paclage and the file would fun.go inside the package
// Now calculate the sum the difference and the product of all five numbers and return it
// you have to call it inside the main.go

func conditionals(param int) {
	// if else else if  switch statments

	if param > 2 {
		fmt.Println("I am greater than 2")
	}
	if param < 2 {
		fmt.Println("I am less than 3")
	}

}

func cond(param int) {
	if param > 2 {
		fmt.Println("hello")
	}
	if param > 4 {
		fmt.Println("hello 2")
	} else {
		fmt.Println("hello 3")
	}
}

// that if we have if else if and else statment the first one of them to be true is executed and compiler ignores others statments
// but if we have multiple if statments in that case all the if statements will be executed and not ignored if they are true

// You have to dd a fizzbuzz problem ---
// You have to crrate a package called as fizzbuzz inside create a file called as fizzbuzz.go and then writ e function in it
// called as Fizzbuzz which will take integera param ,now you have to return string from it
// if the param is divisible by 3  ( num % 3   == 0 )  the return fizz
// if the param is divisble by 5 return buzz
// if the param is dovisble by both 3 and 5 return fizzbuzz

func switchCases(weekDays int) {
	switch weekDays {
	case 1:
		fmt.Println("Its monday")
	case 2:
		fmt.Println("Its tuesday")

	case 3:
		fmt.Println("Its wednesday")
	default:
		fmt.Println("We dnt know what day it is")

	}
}

func fizzBuzz(num int) string {

	switch {
	case num%3 == 0:
		return "fizzz"
	case num%5 == 0:
		return "buzz"
	case num%5 == 0 && num%3 == 0:
		return "fizz buzz"
	default:
		return ""
	}

}

// just replace your logic for the fizzbuzz to have switch case ----->>

func Loop() {

	// first variattio

	//for i := 0; i < 10; i++ {
	//	fmt.Println("hello", i)
	//}
	//
	//// varoation of loop just like while loop
	//
	///**
	//    while condition{
	//
	//    // exprrsiiosn
	//      // condtion change
	//}
	//
	//*/
	//
	//// similar to while loop
	//z := 10
	//for z > 0 {
	//	fmt.Println("hello there")
	//	z-- // z= z-1  // z++
	//}
	//
	//// If you want to have infinite loop
	//
	//for {
	//	fmt.Println("888")
	//
	//	z++
	//
	//	if z > 100 {
	//		break // it will exit from the loop
	//		// continue to exit from the current iteration
	//	}
	//}

	// break vs continue
	//z := 0
	//for {
	//	if z == 78 {
	//		fmt.Println(z)
	//		//continue
	//	}
	//
	//	fmt.Println(z)
	//
	//	if z > 100 {
	//		break
	//	}
	//	z++
	//}
	z := 0
loop:
	for {
		if z > 10 {
			break loop
		}
		z++
	}
}

func evenOdd() {
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue // --->continue
		}
		fmt.Println(i) // will print out all the even numbers now
	}

}

// You have to create a package called as calculator
// you have to defibne function for add subtract and product and division
// then you have to create a command line application
