package tddexercise

// Money
type Money struct {
	amount   int
	currency string
}

func MakeDollar(amount int) Dollar {
	return Dollar{Money{amount: amount, currency: "USD"}}
}

func MakeFranc(amount int) Franc {
	return Franc{Money{amount: amount, currency: "CHF"}}
}

// Dollar
type Dollar struct {
	Money
}

func (d Dollar) times(multiplier int) Money {
	return Money{amount: d.amount * multiplier, currency: d.currency}
}

// Franc
type Franc struct {
	Money
}

func (f Franc) times(multiplier int) Money {
	return Money{amount: f.amount * multiplier, currency: f.currency}
}
