package bankingPkg

type Amount int

func AddAmount(a Amount, b Amount) Amount {
	return a + b
}

func IsPositive(a Amount) bool {
	return a > 0
}

type TransactionStatus int

// statuses
const (
	Pending TransactionStatus = iota 
	Completed                         
	Failed                           
)

// IsFinalStatus returns true if the transaction is Completed or Failed
func IsFinalStatus(status TransactionStatus) bool {
	return status == Completed || status == Failed
}
