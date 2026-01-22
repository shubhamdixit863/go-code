package main

import (
	"fmt"
	"go-code/bankingPkg"
	"go-code/bankingSys"
	"go-code/composition"
	"go-code/directions"
	"go-code/fizbuzz"
	"go-code/funPackage"
	"go-code/highOrder"
	"go-code/ifElse"
	"go-code/interfaces"
	"go-code/loopPrc"
	"go-code/objectsPkg"
	"go-code/errorHandling"
	
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

	//Switch
	nums:=fizbuzz.SwitchPrac(10)
	fmt.Println(nums)

	//continue in LOOPS
	for i:=0;i<50;i++{
		if(i%2==0){
			continue
		}
		fmt.Println(i)
	}
	
	//Loops
	for i := 1; i < 16; i++ {
		fmt.Println(loopPrc.Fizzbuzzes(i))
	}


	var a bankingPkg.Amount = 1500 
	var b bankingPkg.Amount = 2500 

	sum := bankingPkg.AddAmount(a, b)

	fmt.Println("Sum (in paise):", sum)
	fmt.Println("Is sum positive?", bankingPkg.IsPositive(sum))

	// Transaction status
	status := bankingPkg.Completed

	fmt.Println("Is final status?", bankingPkg.IsFinalStatus(status))

	//higher order
	highorder.ApplyOperation(func(n int) int {
		if n%2 == 0 {
			return n * 2 
		} else {
			return n * 3 
	 	}
	})

	//structs+methods
	student1 := objectsPkg.Student{Name: "Rajkumar", Marks: 82}
	student2 := objectsPkg.Student{Name: "Arjun", Marks: 35}

	// Print student details
	fmt.Println("Student Name:", student1.Name)
	fmt.Println("Passed:", student1.IsPassed())
	fmt.Println("Grade:", student1.Grade())
	fmt.Println()

	fmt.Println("Student Name:", student2.Name)
	fmt.Println("Passed:", student2.IsPassed())
	fmt.Println("Grade:", student2.Grade())

	


//ifElse
myAge:=ifElse.VoteEligibility(20)
fmt.Println(myAge)

hi:=userDefined(2,2)
fmt.Println(hi)

result1,result2:=rectangle(2,3)
fmt.Println("Area of rectangle : ",result1,"Perimeter of Rectangle : ",result2)
//Structs
type person struct{
	Name string 
	Age int
	Height int
}
p1:=person{
	Name:"Rajkumar",
	Age: 20,
	Height: 5,
}
fmt.Println(p1.Age,p1.Name,p1.Height)

//Composition 
emp:=composition.Employee{
	Name:"Rajkumar",
	Address:composition.Address{
		City: "Hyderabad",
		State: "Telangana",
	},
}
fmt.Println(emp.Name,
emp.Address.City,
emp.Address.State)

std1:=stud{"Raj",70}
std2:=stud{"other",30}
fmt.Println(std1,stud.IsPassed(std1),stud.Grd(std1))
fmt.Println(std2,stud.IsPassed(std2),stud.Grd(std2))


	/*
	My Banking Aplication 
	*/

	// Customers
	Customer1 := bankingsys.Customer{CustomerID: 1, FullName: "Raam Ayodhya", Age: 21, PhoneNumber: 7671913261, IsActive: true}
	Customer2 := bankingsys.Customer{CustomerID: 2, FullName: "Shiv Kailash", Age: 23, PhoneNumber: 7671913262, IsActive: true}
	Customer3 := bankingsys.Customer{CustomerID: 3, FullName: "Maadhav Dwarakadeesh", Age: 22, PhoneNumber: 7671913263, IsActive: true}

	// Accounts
	Account1 := bankingsys.BankAccount{AccountNumber: 101, AccountType: "SAVINGS", Balance: 5000, IsFrozen: false, Owner: Customer1, LastTxn: bankingsys.Transaction{}}
	Account2 := bankingsys.BankAccount{AccountNumber: 102, AccountType: "CURRENT", Balance: 12000, IsFrozen: false, Owner: Customer2, LastTxn: bankingsys.Transaction{}}
	Account3 := bankingsys.BankAccount{AccountNumber: 103, AccountType: "SAVINGS", Balance: 3000, IsFrozen: false, Owner: Customer3, LastTxn: bankingsys.Transaction{}}

	// Bank
	MyBank := bankingsys.Bank{
		BankName: "Go Bank",
		IFSCCode: "BB0001",
		Account1: Account1,
		Account2: Account2,
		Account3: Account3,
		Customer1: Customer1,
		Customer2: Customer2,
		Customer3: Customer3,
	}

	// Open accounts
	MyBank = MyBank.OpenAccount(Account1)
	MyBank = MyBank.OpenAccount(Account2)
	MyBank = MyBank.OpenAccount(Account3)

	// Operations
	MyBank.Account1 = MyBank.Account1.Deposit(2000, 1)
	MyBank.Account2 = MyBank.Account2.Withdraw(4000, 2)

	// Freeze & fail
	MyBank.Account2 = MyBank.Account2.FreezeAccount()
	MyBank = MyBank.TransferFunds(MyBank.Account2, MyBank.Account1, 2000, 3)

	// Output
	fmt.Println("Account1 Balance:", MyBank.Account1.GetBalance())
	fmt.Println("Last Transaction:", MyBank.Account1.LastTxn)

	fmt.Println("Account2 Balance:", MyBank.Account2.GetBalance())
	fmt.Println("Last Transaction:", MyBank.Account2.LastTxn)

	fmt.Println("Bank Total Funds:", MyBank.GetBankTotalFunds())

	//Interfaces
	mc:=interfaces.McDonalds{
		Menu:"Extra Spicy Burger, French Friess, Sprite",
		FoodPrepResult:"Food Prepared Successfully and will be served within 5 mins !!",
		Location:"Hyderabad",
	}
	kf:=interfaces.KFC{
		Menu:"Fried Chicken from Kentucky, Cheese Pizza, Fanta",
		FoodPrepResult:"Food Preparaton is in progress , Thankyou for your Patience  !!",
		Location:"Bangalore",
	}
	bk:=interfaces.BurgerKing{
		Menu:"chicken Lollipops, Pizza , Thumbs Up",
		FoodPrepResult:"Food Prepared Successfully  Be Ready to eat !!",
		Location:"Karimnagar",
	}
	interfaces.Operate(mc)
	interfaces.Operate(kf)
	interfaces.Operate(bk)
	
	//Area Interface
	rect:=interfaces.Rectangle{
		Width:10.50,
		Height:12.50,
	}
	cir:=interfaces.Circle{
		Radius:25.00,
	}
	fmt.Println("Area of Rectangle is: ",interfaces.CalculateArea(rect))
	fmt.Println("Area of Circle is:",interfaces.CalculateArea(cir))

	//Error Handling
	value, err := errorHandling.Divide(10, 0)
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", value)

}

//User Defined Functions
func userDefined(MyNum1,MyNum2 int)(x int){
	return MyNum1+MyNum2
}

func rectangle(l,w int)(area,perimeter int){
	area=l*w
	perimeter=2*(l+w)
	return
}
/*
 STUDENT USING STRUCT AND METHOD
 */
type stud struct{
	Nam string
	Mks int
}

func (s stud) IsPassed()bool{
	return s.Mks>=40
}
func (s stud) Grd()string{
	if s.Mks>=75{
		return "A"
	}else if s.Mks>=60 && s.Mks<75{
		return "B"
	}else if s.Mks>=40 && s.Mks<60{
		return "C"
	}else{
		return "F"
	}
}

