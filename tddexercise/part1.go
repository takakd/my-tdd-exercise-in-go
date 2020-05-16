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

func (m Money) plus(added Money) Money {
	return Money{amount: m.amount + added.amount, currency: m.currency}
}
