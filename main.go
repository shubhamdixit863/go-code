package main

import (
	"fmt"
	//"go-code/banking"
	"go-code/pointer"
)

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
func main() {

	x := 20
	y := 40

	fmt.Println("Before swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	pointer.SwapAndAdd(&x, &y)

	fmt.Println("\nAfter swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}