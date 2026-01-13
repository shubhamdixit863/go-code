package gobank

type Customer struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
}

func (c *Customer) CanUseBank() bool {
    return c.Age >= 18 && c.IsActive
}
