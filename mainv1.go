package main

import (
	"fmt"
	"reflect"
)

type (
	// ID is user-defined data type
	ID uint32
	// Name is user-defined data type
	Name string
	// Surname is user-defined data type
	Surname string
)

// currency is my own custom type or we can call it as user defined type
type Currency int32

func convert(g Currency) {
	fmt.Println(g)
}

type Human struct {
	Legs      int
	Fingers   int
	HairColor string
}

func NewHuman(legs, finger int, haircolor string) Human {
	return Human{
		Legs:      legs,
		Fingers:   finger,
		HairColor: haircolor,
	}
}

// This is a method
func (h Human) Walk() { // h is  a receiver
	fmt.Println("hello world")
}

type Circle struct {
	Radius int
}

func main() {
	var g Currency
	fmt.Println(reflect.TypeOf(g))
	convert(g)
	//g = 90
	//fmt.Println(g)
	//
	//var name Name
	//name = Name("shubham")

	// type conversion
	g = 9
	fmt.Println(reflect.TypeOf(g))
	// we can do type conversion and we can convert it to float
	h := float32(g) // type conversion
	fmt.Println(reflect.TypeOf(h))

	//eswar := Human{
	//	Legs:      2,
	//	Fingers:   20,
	//	HairColor: "Black",
	//}

	hrithik := NewHuman(2, 21, "blonde")

	fmt.Println(hrithik.Fingers)

	//// anonymous
	//d := struct {
	//	Name string
	//}{
	//	Name: "shubham",
	//}

	//var c Circle
	//c.Radius = 9
	//fmt.Println(c)
	//
	hrithik.Walk()

	//

	h1 := Human{
		Legs:      2,
		Fingers:   201,
		HairColor: "green",
	}

	h2 := NewHuman(2, 20, "green")

	fmt.Println("Equating two structs", h1 == h2) // == is for equality check  !=  equality check operator  = (assignment operator)
	//var c string
	//c="hello"  // we are doing assignment here
	// % --> modulus operator ---> it returns the rmainder of division of two values   5%2 == 1

	PrintHuman(h1)

	s1 := Student{
		Name: "shubham",
		Age:  20,
		Address: Address{
			City:    "New Delhi",
			State:   "New Delhi",
			PinCode: "201001",
		}}

	fmt.Printf("%+v", s1)

	e1 := Employee{
		Address{
			City:    "New Delhi",
			State:   "New Delhi",
			PinCode: "201001",
		},
	}
	fmt.Printf("%+v", e1)
}

// ----go is not an object oriented progrmaming language it has objects
// object is just a real world entity which we represent in our code
/// - how an object --what object contains ---> object has properties and object has behaviour

func PrintHuman(h Human) {
	//fmt.Println(h)

	// you can use formatted printing as well
	//fmt.Printf("The struct %s ,has %d ,%+v", "human", 3, h) // %s and %d are formatter
}

// struct composition
// inheritance which actuall corresponds to is a relationship and there is parentchild relation
// composition corresponds to has a relationship

// If one struct is included in another struct its called composition thats all
type Student struct {
	Name    string
	Age     int
	Address Address // composition ,if you are including (embedding ) one struct into another that becomes composition
}

type Address struct {
	City    string
	State   string
	PinCode string
}

type Employee struct {
	Address // particularly this is embedding
}
