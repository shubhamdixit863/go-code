package bankingsys

type Customer struct {
	CustomerID  int
	FullName    string
	Age         int
	PhoneNumber int64
	IsActive    bool
}

func (c Customer) ActivateCustomer() Customer {
	c.IsActive = true
	return c
}

func (c Customer) DeactivateCustomer() Customer {
	c.IsActive = false
	return c
}

func (c Customer) IsEligibleForBanking() bool {
	return c.Age >= 18 && c.IsActive
}
