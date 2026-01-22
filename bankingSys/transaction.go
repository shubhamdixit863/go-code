package bankingsys

type Transaction struct {
	TransactionID   int
	AccountNumber   int
	TransactionType string
	Amount          float64
	Status          string
}

func CreateTransaction(id int, acc int, t string, amt float64) Transaction {
	return Transaction{
		TransactionID:   id,
		AccountNumber:   acc,
		TransactionType: t,
		Amount:          amt,
	}
}

func (t Transaction) MarkSuccess() Transaction {
	t.Status = "SUCCESS"
	return t
}

func (t Transaction) MarkFailed(reason string) Transaction {
	t.Status = "FAILED"
	return t
}
