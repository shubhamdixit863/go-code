package main

import (
	"fmt"
	"go-code/bankingSystems"
)

func main() {


	cust1 := bankingSystems.Customer{
		CustomerID:  1,
		FullName:    "Arjun",
		Age:         25,
		PhoneNumber: "9999999999",
		IsActive:    true,
	}

	cust2 := bankingSystems.Customer{
		CustomerID:  2,
		FullName:    "Meena",
		Age:         30,
		PhoneNumber: "8888888888",
		IsActive:    true,
	}

	cust3 := bankingSystems.Customer{
		CustomerID:  3,
		FullName:    "Underage User",
		Age:         16,
		PhoneNumber: "7777777777",
		IsActive:    true,
	}


	acc1 := bankingSystems.BankAccount{
		AccountNumber: 101,
		AccountType:   "Savings",
		Balance:       5000,
		Owner:         &cust1,
	}

	acc2 := bankingSystems.BankAccount{
		AccountNumber: 102,
		AccountType:   "Current",
		Balance:       10000,
		Owner:         &cust2,
	}

	acc3 := bankingSystems.BankAccount{
		AccountNumber: 103,
		AccountType:   "Savings",
		Balance:       2000,
		Owner:         &cust3,
	}


	bank := bankingSystems.Bank{
		BankName: "Simple Go Bank",
		IFSCCode: "GO001",
		Account1: acc1,
		Account2: acc2,
		Account3: acc3,
	}


	bank.Account1.Deposit(2000)
	bank.Account2.Withdraw(3000)

	bank.Account1.FreezeAccount()
	bank.Account1.Withdraw(500)

	bank.TransferFunds(&bank.Account2, &bank.Account1, 1500)

	cust2.DeactivateCustomer()
	bank.Account2.Deposit(1000)

	fmt.Println("Account1 Balance:", bank.Account1.GetBalance())
	fmt.Println("Last Txn Status:", bank.Account1.LastTxn.Status)
	fmt.Println("Bank Total Funds:", bank.GetBankTotalFunds())
}
