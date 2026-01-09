package main

import (
	"fmt"
	"go-code/directions"
	"go-code/funPackage"
	"go-code/fizbuzz"
	"go-code/loopPrc"
	
	
)

func main() {

	//Directions using basic Import functions
	fmt.Println("Directions")
	fmt.Println("North: ",directions.North)
	fmt.Println("East: ",directions.East)
	fmt.Println("West: ",directions.West)
	fmt.Println("South: ",directions.South)

	//Arithmatic Operations Using Functions
	Sum,Sub,Mul:=funPackage.Cal(1,2,3,4,5)
	fmt.Println("Sum",Sum)
	fmt.Println("Sub",Sub)
	fmt.Println("Mul",Mul)

	//Division
	div,rem:=funPackage.DivRem(10,20)
	fmt.Println("Division:",div)
	fmt.Println("Remainder:",rem)

	nums:=fizbuzz.SwitchPrac(10)
	fmt.Println(nums)

	for i:=0;i<50;i++{
		if(i%2==0){
			continue
		}
		fmt.Println(i)
	}
	
	for i := 1; i < 16; i++ {
		fmt.Println(loopPrc.Fizzbuzzes(i))
	}


	// ApplyOperation(func(n int) int {
	// 	if n%2 == 0 {
	// 		return n * 2 
	// 	} else {
	// 		return n * 3 
	// // 	}
	// })

}


