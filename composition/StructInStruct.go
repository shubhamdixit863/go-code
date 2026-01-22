package composition

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Address Address
}
