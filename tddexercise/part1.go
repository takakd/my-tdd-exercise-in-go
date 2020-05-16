package tddexercise

// Money
type Money struct {
	amount int
}

func MakeDollar(amount int) Dollar {
	return Dollar{Money{amount: amount}}
}

func MakeFranc(amount int) Franc {
	return Franc{Money{amount: amount}}
}

// Dollar
type Dollar struct {
	Money
}

func (d Dollar) times(multiplier int) Money {
	return MakeDollar(d.amount * multiplier).Money
}

// Franc
type Franc struct {
	Money
}

func (f Franc) times(multiplier int) Money {
	return MakeFranc(f.amount * multiplier).Money
}
