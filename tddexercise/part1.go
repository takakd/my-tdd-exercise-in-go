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

func (m Money) reduce(to string) Money {
	if m.currency == "CHF" && to == "USD" {
		return MakeMoney(m.amount / 2, to)
	}
	return m
}

// Expression
type Expression interface {
	reduce(to string) Money
}

// Bank
type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) reduce(source Expression, to string) Money {
	// bad code.

	// @note: pattern if
	//m, ok := source.(Money)
	//if ok {
	//	return m
	//}
	//sum, ok := source.(*Sum)
	//return sum.reduce(to)

	// @note: pattern type switch
	//switch v := source.(type) {
	//case Money:
	//	return v.reduce(to)
	//case *Sum:
	//	return v.reduce(to)
	//}
	//panic("unknown type")

	return source.reduce(to)
}

func (b *Bank) addRate(from string, to string, rate int) {

}

// Sum
type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend Money, addend Money) *Sum {
	return &Sum{augend: augend, addend: addend}
}

func (s *Sum) reduce(to string) Money {
	amount := s.augend.amount + s.addend.amount
	return MakeMoney(amount, to)
}
