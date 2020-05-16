package tddexercise

// Money
type Money struct {
}

// Dollar
type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}


// Franc
type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (self Franc) times(multiplier int) Franc {
	return NewFranc(self.amount * multiplier)
}
