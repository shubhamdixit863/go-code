package bank

import "fmt"

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

type BankAccount struct {
	AccountNumber int
	AccountType   string // Savings / Current
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
