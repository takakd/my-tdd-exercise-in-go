package tddexercise

// Money
type Money struct {
	amount int
}

// Dollar
type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money{amount: amount}}
}

func (d Dollar) times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}

// Franc
type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{Money{amount: amount}}
}

func (f Franc) times(multiplier int) Franc {
	return NewFranc(f.amount * multiplier)
}
