// package main

// import (
// 	"fmt"
// 	"g0-code/account"
// )

// func main() {

// 	b := account.Bank{
// 		BankName:  "Go National Bank",
// 		IFSCCode: "GONB0001",
// 	}

// 	b.Customer1 = account.Customer{CustomerID: 1, FullName: "Eswar", Age: 23, IsActive: true}
// 	b.Customer2 = account.Customer{CustomerID: 2, FullName: "Mahesh", Age: 22, IsActive: true}
// 	b.Customer3 = account.Customer{CustomerID: 3, FullName: "Sai krishna", Age: 23, IsActive: true}

// 	b.OpenAccount(&b.Customer1, 101, "Savings", 2000)
// 	b.OpenAccount(&b.Customer2, 102, "Current", 10000)
// 	b.OpenAccount(&b.Customer3, 103, "Savings", 500)

// 	b.Customer1.Account.Deposit(500, b.Customer1.IsActive)
// 	b.Customer1.Account.Withdraw(1000, b.Customer1.IsActive)

// 	b.TransferFunds(&b.Customer1, &b.Customer2, 300)

// 	b.Customer1.Account.FreezeAccount()
// 	b.Customer1.Account.Withdraw(100, b.Customer1.IsActive)

// 	b.Customer2.DeactivateCustomer()
// 	b.Customer2.Account.Deposit(1000, b.Customer2.IsActive)

// 	fmt.Println("\n--- FINAL STATE ---")
// 	fmt.Println("Customer1 Balance:", b.Customer1.Account.GetBalance())
// 	fmt.Println("Customer2 Balance:", b.Customer2.Account.GetBalance())
// 	fmt.Println("Bank Total Funds:", b.GetBankTotalFunds())
// }


package main

import (
	"fmt"
	"g0-code/interfaces"
)


func Operate(fr interfaces.Franchise) {
	fmt.Println("Location:", fr.GetLocation())
	fmt.Println("Menu:")
	for _, item := range fr.GetMenu() {
		fmt.Println("-", item)
	}
	fmt.Println("Operation:", fr.PrepareFood())
	
}

func main() {
	mc := interfaces.McDonalds{}
	kfc := interfaces.KFC{}
	bk := interfaces.BurgerKing{}

	Operate(mc)
	Operate(kfc)
	Operate(bk)
}