package tddexercise

// Money
type Money struct {
	amount int
}

func MakeDollar(amount int) Money {
	return Dollar{Money{amount: amount}}.Money
}

func MakeFranc(amount int) Money {
	return Franc{Money{amount: amount}}.Money
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
