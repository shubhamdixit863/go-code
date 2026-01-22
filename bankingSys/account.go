package bankingsys

type BankAccount struct {
	AccountNumber int
	AccountType   string
	Balance       float64
	IsFrozen      bool
	Owner         Customer
	LastTxn       Transaction
}

func (a BankAccount) Deposit(amount float64, txnID int) BankAccount {
	txn := CreateTransaction(txnID, a.AccountNumber, "DEPOSIT", amount)

	if !a.Owner.IsEligibleForBanking() {
		a.LastTxn = txn.MarkFailed("Customer not eligible")
		return a
	}

	a.Balance += amount
	a.LastTxn = txn.MarkSuccess()
	return a
}

func (a BankAccount) Withdraw(amount float64, txnID int) BankAccount {
	txn := CreateTransaction(txnID, a.AccountNumber, "WITHDRAW", amount)

	if a.IsFrozen {
		a.LastTxn = txn.MarkFailed("Account is frozen")
		return a
	}

	minBalance := 1000.0
	if a.AccountType == "CURRENT" {
		minBalance = 5000.0
	}

	if a.Balance-amount < minBalance {
		a.LastTxn = txn.MarkFailed("Minimum balance violation")
		return a
	}

	a.Balance -= amount
	a.LastTxn = txn.MarkSuccess()
	return a
}

func (a BankAccount) GetBalance() float64 {
	return a.Balance
}

func (a BankAccount) FreezeAccount() BankAccount {
	a.IsFrozen = true
	return a
}

func (a BankAccount) UnfreezeAccount() BankAccount {
	a.IsFrozen = false
	return a
}
