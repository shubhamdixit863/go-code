package main

import (
	"fmt"
	"go-code/banking"
	"go-code/franchise"
)

func main() {

	// 1. Create Bank
	bank := banking.Bank{
		BankName: "Go Bank",
		IFSCCode: "GO001",
	}

	// 2. Create Customers
	c1 := banking.Customer{1, "Amit", 25, "9999", true}
	c2 := banking.Customer{2, "Sneha", 30, "8888", true}
	c3 := banking.Customer{3, "Ravi", 17, "7777", true}

	// 3. Open Accounts
	bank.Account1 = banking.BankAccount{101, "Savings", 2000, false, false, &c1, banking.Transaction{}}
	bank.Account2 = banking.BankAccount{102, "Current", 8000, false, false, &c2, banking.Transaction{}}
	bank.Account3 = banking.BankAccount{103, "Savings", 1500, false, false, &c3, banking.Transaction{}}

	// 4. Perform operations
	bank.Account1.Deposit(500)
	bank.Account1.Withdraw(200)

	// 5. Freeze account and show failure
	bank.Account1.FreezeAccount()
	bank.Account1.Withdraw(100)

	// 6. Deactivate customer and block transaction
	c2.DeactivateCustomer()
	bank.Account2.Deposit(100)

	// 7. Transfer funds
	bank.TransferFunds(&bank.Account2, &bank.Account1, 500)

	// 8. Print results
	fmt.Println("\n--- FINAL REPORT ---")
	fmt.Println("Account 101 Balance:", bank.Account1.GetBalance())
	fmt.Println("Last Txn Status:", bank.Account1.LastTxn.Status)

	fmt.Println("Account 102 Balance:", bank.Account2.GetBalance())
	fmt.Println("Last Txn Status:", bank.Account2.LastTxn.Status)

	fmt.Println("Bank Total Funds:", bank.GetBankTotalFunds())


    mc:=franchise.McDonalds{
		Location: "Chennai",
		Menu: "Burger, Fries, Coke",
		Preparation: "Burger is prepared with fresh ingredients.",
	}
	kfc:=franchise.KFC{
		Location: "Bangalore",
		Menu: "Fried Chicken, Mashed Potatoes, Gravy",
		Preparation: "Fried Chicken is prepared with secret spices.",
	}
	bk:=franchise.BurgerKing{
		Location: "Mumbai",
		Menu: "Whopper, Onion Rings, Soft Drink",
		Preparation: "Whopper is flame-grilled to perfection.",
	}
	fmt.Println("\n--- FRANCHISE OPERATIONS ---")
	franchise.Operate(mc)
	franchise.Operate(kfc)
	franchise.Operate(bk)

}