package account

import "fmt"



type Transaction struct {
	TransactionID   int
	AccountNumber   int
	TransactionType string
	Amount          float64
	Status          string
	FailureReason   string
}

func (t *Transaction) MarkSuccess() {
	t.Status = "SUCCESS"
	t.FailureReason = ""
}

func (t *Transaction) MarkFailed(reason string) {
	t.Status = "FAILED"
	t.FailureReason = reason
}



type BankAccount struct {
	AccountNumber   int
	AccountType     string
	Balance         float64
	IsFrozen        bool
	LastTransaction Transaction
}

func (a *BankAccount) minBalance() float64 {
	if a.AccountType == "Savings" {
		return 1000
	}
	return 5000
}

func (a *BankAccount) Deposit(amount float64, customerActive bool) {
	t := Transaction{
		TransactionID:   a.LastTransaction.TransactionID + 1,
		AccountNumber:   a.AccountNumber,
		TransactionType: "DEPOSIT",
		Amount:          amount,
	}

	if !customerActive {
		t.MarkFailed("Customer is inactive")
		a.LastTransaction = t
		fmt.Println("Deposit failed:", t.FailureReason)
		return
	}

	if a.IsFrozen {
		t.MarkFailed("Account is frozen")
		a.LastTransaction = t
		fmt.Println("Deposit failed:", t.FailureReason)
		return
	}

	a.Balance += amount
	t.MarkSuccess()
	a.LastTransaction = t
	fmt.Println("Deposit successful")
}

func (a *BankAccount) Withdraw(amount float64, customerActive bool) bool {
	t := Transaction{
		TransactionID:   a.LastTransaction.TransactionID + 1,
		AccountNumber:   a.AccountNumber,
		TransactionType: "WITHDRAW",
		Amount:          amount,
	}

	if !customerActive {
		t.MarkFailed("Customer is inactive")
		a.LastTransaction = t
		fmt.Println("Withdraw failed:", t.FailureReason)
		return false
	}

	if a.IsFrozen {
		t.MarkFailed("Account is frozen")
		a.LastTransaction = t
		fmt.Println("Withdraw failed:", t.FailureReason)
		return false
	}

	if a.Balance-amount < a.minBalance() {
		t.MarkFailed("Minimum balance violation")
		a.LastTransaction = t
		fmt.Println("Withdraw failed:", t.FailureReason)
		return false
	}

	a.Balance -= amount
	t.MarkSuccess()
	a.LastTransaction = t
	fmt.Println("Withdraw successful")
	return true
}

func (a *BankAccount) FreezeAccount() {
	a.IsFrozen = true
}

func (a *BankAccount) UnfreezeAccount() {
	a.IsFrozen = false
}

func (a *BankAccount) GetBalance() float64 {
	return a.Balance
}



type Customer struct {
	CustomerID int
	FullName   string
	Age        int
	IsActive   bool
	Account    BankAccount
}

func (c *Customer) ActivateCustomer() {
	c.IsActive = true
}

func (c *Customer) DeactivateCustomer() {
	c.IsActive = false
}

func (c *Customer) IsEligibleForBanking() bool {
	return c.Age >= 18 && c.IsActive
}



type Bank struct {
	BankName  string
	IFSCCode string
	TotalFund float64

	Customer1 Customer
	Customer2 Customer
	Customer3 Customer
}

func (b *Bank) OpenAccount(c *Customer, accNo int, accType string, initialBalance float64) {
	if c.Age < 18 {
		fmt.Println("Account opening failed: Underage customer")
		return
	}

	min := 1000.0
	if accType == "Current" {
		min = 5000
	}

	if initialBalance < min {
		fmt.Println("Account opening failed: Minimum balance not met")
		return
	}

	c.Account = BankAccount{
		AccountNumber: accNo,
		AccountType:   accType,
		Balance:       initialBalance,
	}

	b.TotalFund += initialBalance
	fmt.Println("Account opened for", c.FullName)
}

func (b *Bank) TransferFunds(from *Customer, to *Customer, amount float64) {
	fmt.Println("Initiating transfer...")

	if !from.Account.Withdraw(amount, from.IsActive) {
		fmt.Println("Transfer failed")
		return
	}

	to.Account.Deposit(amount, to.IsActive)
	fmt.Println("Transfer successful")
}

func (b *Bank) GetBankTotalFunds() float64 {
	return b.TotalFund
}
