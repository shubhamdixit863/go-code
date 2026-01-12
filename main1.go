package main

import "fmt"



type AccountType int
type TransactionType int
type TransactionStatus int

const (
	Savings AccountType = iota
	Current
)

const (
	Deposit TransactionType = iota
	Withdraw
	Transfer
)

const (
	Success TransactionStatus = iota
	Failed
)



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



type Transaction struct {
	TransactionID  int
	AccountNumber  int
	Type           TransactionType
	Amount         int
	Status         TransactionStatus
	Reason         string
}

func (t *Transaction) MarkSuccess() {
	t.Status = Success
}

func (t *Transaction) MarkFailed(reason string) {
	t.Status = Failed
	t.Reason = reason
	fmt.Println("Transaction Failed:", reason)
}



type BankAccount struct {
	AccountNumber int
	AccountType   AccountType
	Balance       int
	IsFrozen      bool
	Owner         *Customer
	LastTxn       Transaction
}

func (a BankAccount) minimumBalance() int {
	if a.AccountType == Savings {
		return 1000
	}
	return 5000
}

func (a *BankAccount) Deposit(amount int) {
	a.LastTxn = Transaction{1, a.AccountNumber, Deposit, amount, Failed, ""}

	if !a.Owner.IsEligibleForBanking() {
		a.LastTxn.MarkFailed("Customer not eligible")
		return
	}

	a.Balance += amount
	a.LastTxn.MarkSuccess()
}

func (a *BankAccount) Withdraw(amount int) {
	a.LastTxn = Transaction{2, a.AccountNumber, Withdraw, amount, Failed, ""}

	if a.IsFrozen {
		a.LastTxn.MarkFailed("Account is frozen")
		return
	}

	if !a.Owner.IsEligibleForBanking() {
		a.LastTxn.MarkFailed("Customer inactive")
		return
	}

	if a.Balance-amount < a.minimumBalance() {
		a.LastTxn.MarkFailed("Minimum balance violation")
		return
	}

	a.Balance -= amount
	a.LastTxn.MarkSuccess()
}

func (a *BankAccount) GetBalance() int {
	return a.Balance
}

func (a *BankAccount) FreezeAccount() {
	a.IsFrozen = true
}

func (a *BankAccount) UnfreezeAccount() {
	a.IsFrozen = false
}



type Bank struct {
	BankName  string
	IFSCCode string

	Customer1 Customer
	Customer2 Customer
	Customer3 Customer

	Account1 BankAccount
	Account2 BankAccount
	Account3 BankAccount
}

func (b *Bank) OpenAccount(c *Customer, a *BankAccount) {
	if c.Age < 18 {
		fmt.Println("Underage customer cannot open account")
		return
	}
	c.ActivateCustomer()
	fmt.Println("Account opened for:", c.FullName)
}

func (b *Bank) TransferFunds(from *BankAccount, to *BankAccount, amount int) {

	from.LastTxn = Transaction{3, from.AccountNumber, Transfer, amount, Failed, ""}

	if from.IsFrozen {
		from.LastTxn.MarkFailed("Source account frozen")
		return
	}

	if from.Balance-amount < from.minimumBalance() {
		from.LastTxn.MarkFailed("Insufficient balance")
		return
	}

	from.Balance -= amount
	to.Balance += amount
	from.LastTxn.MarkSuccess()
}

func (b Bank) GetBankTotalFunds() int {
	return b.Account1.Balance + b.Account2.Balance + b.Account3.Balance
}



func main() {

	bank := Bank{
		BankName:  "Go Bank",
		IFSCCode: "GOBK0001",
	}

	// Customers
	bank.Customer1 = Customer{1, "Ravi", 30, "9999", false}
	bank.Customer2 = Customer{2, "Anita", 25, "8888", false}
	bank.Customer3 = Customer{3, "Rahul", 17, "7777", false}

	// Accounts
	bank.Account1 = BankAccount{101, Savings, 2000, false, &bank.Customer1, Transaction{}}
	bank.Account2 = BankAccount{102, Current, 6000, false, &bank.Customer2, Transaction{}}
	bank.Account3 = BankAccount{103, Savings, 500, false, &bank.Customer3, Transaction{}}

	// Open accounts
	bank.OpenAccount(&bank.Customer1, &bank.Account1)
	bank.OpenAccount(&bank.Customer2, &bank.Account2)
	bank.OpenAccount(&bank.Customer3, &bank.Account3) // 
	// Transactions
	bank.Account1.Deposit(1000)
	bank.Account2.Withdraw(2000)

	// Freeze account
	bank.Account2.FreezeAccount()
	bank.Account2.Withdraw(500)

	// Transfer
	bank.TransferFunds(&bank.Account1, &bank.Account2, 500)

	// Deactivate customer
	bank.Customer1.DeactivateCustomer()
	bank.Account1.Withdraw(200)

	// Final output
	fmt.Println("\nFINAL SUMMARY")
	fmt.Println("Account1 Balance:", bank.Account1.GetBalance())
	fmt.Println("Account2 Balance:", bank.Account2.GetBalance())
	fmt.Println("Bank Total Funds:", bank.GetBankTotalFunds())
}
