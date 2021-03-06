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

func (m Money) times(multiplier int) Expression {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) plus(added Expression) Expression {
	return NewSum(m, added)
}

func (m Money) reduce(bank *Bank, to string) Money {
	rate := bank.rate(m.currency, to)
	return MakeMoney(m.amount/rate, to)
}

// Expression
type Expression interface {
	plus(added Expression) Expression
	reduce(bank *Bank, to string) Money
	times(multiplier int) Expression
}

// Bank
type Bank struct {
	rates map[struct {
		from string
		to   string
	}]int
}

func NewBank() *Bank {
	return &Bank{
		rates: map[struct {
			from string
			to   string
		}]int{}}
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

	return source.reduce(b, to)
}

func (b *Bank) addRate(from string, to string, rate int) {
	b.rates[struct {
		from string
		to   string
	}{from, to}] = rate
}

func (b *Bank) rate(from string, to string) int {
	if from == to {
		return 1
	}

	rate, ok := b.rates[struct {
		from string
		to   string
	}{from, to}]
	if ! ok {
		panic("unknown key")
	}
	return rate
}

// Sum
type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) *Sum {
	return &Sum{augend: augend, addend: addend}
}

func (s *Sum) reduce(bank *Bank, to string) Money {
	amount := s.augend.reduce(bank, to).amount + s.addend.reduce(bank, to).amount
	return MakeMoney(amount, to)
}

func (s *Sum) plus(added Expression) Expression {
	return NewSum(s, added)
}

func (s *Sum) times(multiplier int) Expression {
	return NewSum(s.augend.times(multiplier), s.addend.times(multiplier))
}
