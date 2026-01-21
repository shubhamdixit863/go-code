//package main

// import (
// 	"fmt"
// 	//"go-code/banking"
// 	//"go-code/pointer"
// 	"go-code/restaurant"
// )

// func main() {


// 	bank := banking.Bank{
// 		BankName: "Go Bank",
// 		IFSCCode: "GO001",
// 	}

	
// 	c1 := banking.Customer{1, "Amit", 25, "9999", true}
// 	c2 := banking.Customer{2, "Sneha", 30, "8888", true}
// 	c3 := banking.Customer{3, "Ravi", 17, "7777", true}

	
// 	bank.Account1 = banking.BankAccount{101, "Savings", 2000, false, false, &c1, banking.Transaction{}}
// 	bank.Account2 = banking.BankAccount{102, "Current", 8000, false, false, &c2, banking.Transaction{}}
// 	bank.Account3 = banking.BankAccount{103, "Savings", 1500, false, false, &c3, banking.Transaction{}}

	
// 	bank.Account1.Deposit(500)
// 	bank.Account1.Withdraw(200)

	
// 	bank.Account1.FreezeAccount()
// 	bank.Account1.Withdraw(100)

	
// 	c2.DeactivateCustomer()
// 	bank.Account2.Deposit(100)

	
// 	bank.TransferFunds(&bank.Account2, &bank.Account1, 500)

	
// 	fmt.Println("\n--- FINAL REPORT ---")
// 	fmt.Println("Account 101 Balance:", bank.Account1.GetBalance())
// 	fmt.Println("Last Txn Status:", bank.Account1.LastTxn.Status)

// 	fmt.Println("Account 102 Balance:", bank.Account2.GetBalance())
// 	fmt.Println("Last Txn Status:", bank.Account2.LastTxn.Status)

// 	fmt.Println("Bank Total Funds:", bank.GetBankTotalFunds())
// }
// func main() {

// 	x := 20
// 	y := 40

// 	fmt.Println("Before swap:")
// 	fmt.Println("x =", x)
// 	fmt.Println("y =", y)
// 	pointer.SwapAndAdd(&x, &y)

// 	fmt.Println("\nAfter swap:")
// 	fmt.Println("x =", x)
// 	fmt.Println("y =", y)
// }
package main

import "fmt"

type Franchise interface {
	PrepareFood() string
	GetMenu() []string
	GetLocation() string
}

type McDonalds struct{}

func (McDonalds) PrepareFood() string { return "Big Mac ready" }
func (McDonalds) GetMenu() []string   { return []string{"Big Mac", "Fries"} }
func (McDonalds) GetLocation() string { return "Hyderabad" }

type KFC struct{}

func (KFC) PrepareFood() string { return "Chicken ready" }
func (KFC) GetMenu() []string   { return []string{"Fried Chicken", "Zinger"} }
func (KFC) GetLocation() string { return "Bangalore" }

type BurgerKing struct{}

func (BurgerKing) PrepareFood() string { return "Whopper ready" }
func (BurgerKing) GetMenu() []string   { return []string{"Whopper", "Fries"} }
func (BurgerKing) GetLocation() string { return "Chennai" }

func Operate(f Franchise) {
	fmt.Println(f.GetLocation())
	fmt.Println(f.GetMenu())
	fmt.Println(f.PrepareFood())
}

func main() {
	Operate(McDonalds{})
	Operate(KFC{})
	Operate(BurgerKing{})
}
