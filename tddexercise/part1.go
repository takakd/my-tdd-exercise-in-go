package tddexercise

// Money
type Money struct {
	amount   int
	currency string
}

func MakeDollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func MakeFranc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}

func (m Money) times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) plus(added Money) Expression {
	return Money{amount: m.amount + added.amount, currency: m.currency}
}

// Expression
type Expression interface {
}

// Bank
type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) reduce(source Expression, to string) Money {
	return MakeDollar(10)
}

// Sum
type Sum struct {
	augend int
	addend int
}
