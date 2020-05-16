package tddexercise

// Money
type Money struct {
	amount   int
	currency string
}

func MakeMoney(amount int, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func MakeDollar(amount int) Money {
	return MakeMoney(amount, "USD")
}

func MakeFranc(amount int) Money {
	return MakeMoney(amount, "CHF")
}

func (m Money) times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) plus(added Money) Expression {
	return NewSum(m, added)
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
	sum := source.(*Sum)
	return sum.reduce(to)
}

// Sum
type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend Money, addend Money) *Sum {
	return &Sum{augend: augend, addend: addend}
}

func (s *Sum)reduce(to string) Money {
	amount := s.augend.amount + s.addend.amount
	return MakeMoney(amount, to)
}
