package bankingsys

type Bank struct {
	BankName   string
	IFSCCode   string
	TotalFunds float64

	Customer1 Customer
	Customer2 Customer
	Customer3 Customer

	Account1 BankAccount
	Account2 BankAccount
	Account3 BankAccount
}

func (b Bank) OpenAccount(acc BankAccount) Bank {
	b.TotalFunds += acc.Balance
	return b.updateAccount(acc)
}

func (b Bank) TransferFunds(from BankAccount, to BankAccount, amount float64, txnID int) Bank {
	txn := CreateTransaction(txnID, from.AccountNumber, "TRANSFER", amount)

	if from.IsFrozen {
		from.LastTxn = txn.MarkFailed("Source account frozen")
		return b.updateAccount(from)
	}

	if from.Balance < amount {
		from.LastTxn = txn.MarkFailed("Insufficient balance")
		return b.updateAccount(from)
	}

	from.Balance -= amount
	to.Balance += amount

	from.LastTxn = txn.MarkSuccess()

	b = b.updateAccount(from)
	b = b.updateAccount(to)

	return b
}

func (b Bank) GetBankTotalFunds() float64 {
	return b.TotalFunds
}

func (b Bank) updateAccount(acc BankAccount) Bank {
	if acc.AccountNumber == b.Account1.AccountNumber {
		b.Account1 = acc
	} else if acc.AccountNumber == b.Account2.AccountNumber {
		b.Account2 = acc
	} else if acc.AccountNumber == b.Account3.AccountNumber {
		b.Account3 = acc
	}
	return b
}
