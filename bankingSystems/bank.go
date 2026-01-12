package bankingSystems

import "fmt"

/* ---------------- TRANSACTION ---------------- */

type Transaction struct {
	TransactionID   int
	AccountNumber   int
	TransactionType string
	Amount          float64
	Status          string
	Reason          string
}

func (t *Transaction) MarkSuccess() {
	t.Status = "SUCCESS"
	t.Reason = ""
}

func (t *Transaction) MarkFailed(reason string) {
	t.Status = "FAILED"
	t.Reason = reason
	fmt.Println("Transaction Failed:", reason)
}

/* ---------------- CUSTOMER ---------------- */

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

func (c *Customer) IsEligibleForBanking() bool {
	return c.Age >= 18 && c.IsActive
}



type BankAccount struct {
	AccountNumber int
	AccountType   string
	Balance       float64
	IsFrozen      bool
	IsClosed      bool
	Owner         *Customer
	LastTxn       Transaction
}



func (a *BankAccount) Deposit(amount float64) {
	txn := Transaction{1, a.AccountNumber, "DEPOSIT", amount, "", ""}

	if a.IsClosed == true {
		txn.MarkFailed("Deposit into closed account")
		a.LastTxn = txn
		return
	}

	eligible := a.Owner.IsEligibleForBanking()
	if eligible == false {
		txn.MarkFailed("Customer inactive or underage")
		a.LastTxn = txn
		return
	}

	a.Balance += amount
	txn.MarkSuccess()
	a.LastTxn = txn
}


func (a *BankAccount) Withdraw(amount float64) {
	txn := Transaction{2, a.AccountNumber, "WITHDRAW", amount, "", ""}

	if a.IsFrozen == true {
		txn.MarkFailed("Account is frozen")
		a.LastTxn = txn
		return
	}

	minBalance := 1000.0
	if a.AccountType == "Current" {
		minBalance = 5000
	}

	newBalance := a.Balance - amount
	if newBalance < minBalance {
		txn.MarkFailed("Minimum balance violation")
		a.LastTxn = txn
		return
	}

	a.Balance = newBalance
	txn.MarkSuccess()
	a.LastTxn = txn
}



func (a *BankAccount) GetBalance() float64 {
	return a.Balance
}

func (a *BankAccount) FreezeAccount() {
	a.IsFrozen = true
}

func (a *BankAccount) UnfreezeAccount() {
	a.IsFrozen = false
}

func (a *BankAccount) CloseAccount() {
	a.IsClosed = true
}



type Bank struct {
	BankName string
	IFSCCode string
	Account1 BankAccount
	Account2 BankAccount
	Account3 BankAccount
}



func (b *Bank) TransferFunds(from *BankAccount, to *BankAccount, amount float64) {
	txn := Transaction{3, from.AccountNumber, "TRANSFER", amount, "", ""}

	if from.IsFrozen == true {
		txn.MarkFailed("Transfer from frozen account")
		from.LastTxn = txn
		return
	}

	hasBalance := from.Balance >= amount
	if hasBalance == false {
		txn.MarkFailed("Insufficient balance")
		from.LastTxn = txn
		return
	}

	from.Balance -= amount
	to.Balance += amount
	txn.MarkSuccess()
	from.LastTxn = txn
}



func (b *Bank) GetBankTotalFunds() float64 {
	return b.Account1.Balance +
		b.Account2.Balance +
		b.Account3.Balance
}
