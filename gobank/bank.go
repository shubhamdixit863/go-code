package gobank

type Bank struct {
    Name string

    Acc1 Account
    Acc2 Account
    Acc3 Account
}

func (b *Bank) OpenAccount(c *Customer, accNo int, accType string, money float64) *Account {

    acc := Account{
        Number:  accNo,
        Type:    accType,
        Balance: money,
        Owner:   c,
    }

    if b.Acc1.Number == 0 {
        b.Acc1 = acc
        return &b.Acc1
    }

    if b.Acc2.Number == 0 {
        b.Acc2 = acc
        return &b.Acc2
    }

    b.Acc3 = acc
    return &b.Acc3
}

func (b *Bank) Transfer(from, to *Account, amt float64) Transaction {

    w := from.Withdraw(amt)
    if w.Status == "Failed" {
        return w
    }

    d := to.Deposit(amt)
    return d
}
