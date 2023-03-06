package interfaces

type withdrawer interface {
	Withdraw(amount float64) error
}

type depositor interface {
	Deposit(amount float64)
}

type balancer interface {
	Balance() float64
}

type Banker interface {
	balancer
	withdrawer
	depositor
}
