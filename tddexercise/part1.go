package tddexercise

// Money
type Money struct {
	amount int
}

func (m Money)times(multipiler int) Money {
	return Money{}
}

func (m Money)currency() string {
	return ""
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
	return MakeDollar(d.amount * multiplier)
}

func (d Dollar) currency() string {
	return "USD"
}

// Franc
type Franc struct {
	Money
}

func (f Franc) times(multiplier int) Money {
	return MakeFranc(f.amount * multiplier)
}

func (f Franc) currency() string {
	return "CHF"
}
