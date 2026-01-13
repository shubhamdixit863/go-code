package gobank

type Account struct {
    Number int
    Type   string
    Balance float64
    Owner   *Customer
    LastTxn Transaction
}

func (a *Account) Deposit(amount float64) Transaction {
    t := Transaction{Type: "Deposit", Amount: amount}

    if !a.Owner.CanUseBank() {
        t.Status = "Failed"
        t.Note = "Customer not eligible"
        a.LastTxn = t
        return t
    }

    a.Balance += amount
    t.Status = "Success"
    a.LastTxn = t
    return t
}

func (a *Account) Withdraw(amount float64) Transaction {
    t := Transaction{Type: "Withdraw", Amount: amount}

    if !a.Owner.CanUseBank() {
        t.Status = "Failed"
        t.Note = "Customer not eligible"
        a.LastTxn = t
        return t
    }

    if a.Balance < amount {
        t.Status = "Failed"
        t.Note = "Not enough balance"
        a.LastTxn = t
        return t
    }

    a.Balance -= amount
    t.Status = "Success"
    a.LastTxn = t
    return t
}
