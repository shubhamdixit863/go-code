package main

import "fmt"

//transaction

type Transaction struct {
	TransactionID   int
	AccountNumber   int
	TransactionType string
	Amount          int
	Status          string
	Reason          string
}

func (t *Transaction) MarkSuccess() {
	t.Status = "SUCCESS"
}

func (t *Transaction) MarkFailed(reason string) {
	t.Status = "FAILED"
	t.Reason = reason
	fmt.Println("Transaction Failed:", reason)
}

//CUSTOMER

type Customer struct {
	CustomerID  int
	FullName    string
	Age         int
	PhoneNumber string
	IsActive    bool
}

func (c *Customer) ActivateCustomer() {
	c.IsActive = true
}

func (c *Customer) DeactivateCustomer() {
	c.IsActive = false
}

func (c Customer) IsEligibleForBanking() bool {
	return c.Age >= 18 && c.IsActive
}

//BANK ACCOUNT

type BankAccount struct {
	AccountNumber int
	AccountType   string
	Balance       int
	IsFrozen      bool
	LastTxn       Transaction
	Owner         *Customer
}

func (a *BankAccount) minBalance() int {
	if a.AccountType == "Savings" {
		return 1000
	}
	return 5000
}

func (a *BankAccount) Deposit(amount int) {
	a.LastTxn = Transaction{TransactionType: "Deposit", Amount: amount}

	if !a.Owner.IsEligibleForBanking() {
		a.LastTxn.MarkFailed("Customer not eligible")
		return
	}

	a.Balance += amount
	a.LastTxn.MarkSuccess()
}

func (a *BankAccount) Withdraw(amount int) {
	a.LastTxn = Transaction{TransactionType: "Withdraw", Amount: amount}

	if a.IsFrozen {
		a.LastTxn.MarkFailed("Account is frozen")
		return
	}

	if !a.Owner.IsEligibleForBanking() {
		a.LastTxn.MarkFailed("Customer inactive")
		return
	}

	if a.Balance-amount < a.minBalance() {
		a.LastTxn.MarkFailed("Minimum balance violation")
		return
	}

	a.Balance -= amount
	a.LastTxn.MarkSuccess()
}

func (a *BankAccount) FreezeAccount() {
	a.IsFrozen = true
}

func (a *BankAccount) UnfreezeAccount() {
	a.IsFrozen = false
}

func (a BankAccount) GetBalance() int {
	return a.Balance
}

// BANK

type Bank struct {
	BankName string
	IFSCCode string
	Acc1     BankAccount
	Acc2     BankAccount
	Acc3     BankAccount
}

func (b *Bank) TransferFunds(from *BankAccount, to *BankAccount, amount int) {
	from.LastTxn = Transaction{TransactionType: "Transfer", Amount: amount}

	if from.IsFrozen || to.IsFrozen {
		from.LastTxn.MarkFailed("Account frozen")
		return
	}

	if from.Balance-amount < from.minBalance() {
		from.LastTxn.MarkFailed("Insufficient balance")
		return
	}

	from.Balance -= amount
	to.Balance += amount
	from.LastTxn.MarkSuccess()
}

func (b Bank) GetBankTotalFunds() int {
	return b.Acc1.Balance + b.Acc2.Balance + b.Acc3.Balance
}

func main() {
	c1 := Customer{1, "Mahesh", 22, "9999999999", true}
	c2 := Customer{2, "Ravi", 30, "8888888888", true}
	c3 := Customer{3, "Suresh", 17, "7777777777", true}

	acc1 := BankAccount{101, "Savings", 2000, false, Transaction{}, &c1}
	acc2 := BankAccount{102, "Current", 6000, false, Transaction{}, &c2}
	acc3 := BankAccount{103, "Savings", 1000, false, Transaction{}, &c3}

	bank := Bank{"Go Bank", "GO001", acc1, acc2, acc3}

	bank.Acc1.Deposit(500)
	bank.Acc2.Withdraw(1000)
	bank.TransferFunds(&bank.Acc1, &bank.Acc2, 500)

	bank.Acc1.FreezeAccount()
	bank.Acc1.Withdraw(200)

	bank.Acc3.Deposit(100)

	c2.DeactivateCustomer()
	bank.Acc2.Withdraw(500)

	fmt.Println("\nFinal Balances:")
	fmt.Println(bank.Acc1.GetBalance())
	fmt.Println(bank.Acc2.GetBalance())
	fmt.Println(bank.Acc3.GetBalance())

	fmt.Println("Bank Total Funds:", bank.GetBankTotalFunds())
}
