package tddexercise

// Money
type Money struct {
	amount   int
	currency string
}

func MakeDollar(amount int) Money {
	return Dollar{Money{amount: amount, currency: "USD"}}.Money
}

func MakeFranc(amount int) Money {
	return Franc{Money{amount: amount, currency: "CHF"}}.Money
}

func (m Money) times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

// Dollar
type Dollar struct {
	Money
}

// Franc
type Franc struct {
	Money
}
