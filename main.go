package main

import (
    "fmt"
    "gotask/gobank"
)

func main() {

    bank := gobank.Bank{Name: "Simple Bank"}

    // Customers
    c1 := gobank.Customer{ID: 1, Name: "Alice", Age: 25, IsActive: true}
    c2 := gobank.Customer{ID: 2, Name: "Bob", Age: 17, IsActive: true}
    c3 := gobank.Customer{ID: 3, Name: "Charlie", Age: 30, IsActive: true}

    // Open accounts
    a1 := bank.OpenAccount(&c1, 101, "Savings", 5000)
    a2 := bank.OpenAccount(&c2, 102, "Savings", 2000)
    a3 := bank.OpenAccount(&c3, 103, "Current", 7000)

    fmt.Println("Deposit into A1:", a1.Deposit(2000))
    fmt.Println("Deposit into A2:", a2.Deposit(500))  // FAIL

    fmt.Println("Withdraw from A1:", a1.Withdraw(3000))
    fmt.Println("Withdraw from A3:", a3.Withdraw(2000))

    fmt.Println("Transfer A1 → A3:", bank.Transfer(a1, a3, 1000))

    fmt.Println("\nFinal Balances:")
    fmt.Println("A1:", a1.Balance, ", Last:", a1.LastTxn)
    fmt.Println("A2:", a2.Balance, ", Last:", a2.LastTxn)
    fmt.Println("A3:", a3.Balance, ", Last:", a3.LastTxn)
}
