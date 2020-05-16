package tddexercise

// Money
type Money interface {
	amount() int
	times(multipiler int) Money
	currency() string
}

func MakeDollar(amount int) Money {
	return Dollar{amountValue: amount}
}

func MakeFranc(amount int) Money {
	return Franc{amountValue: amount}
}

// Dollar
type Dollar struct {
	amountValue int
}

func (d Dollar) amount() int {
	return d.amountValue
}

func (d Dollar) times(multiplier int) Money {
	return MakeDollar(d.amountValue * multiplier)
}

func (d Dollar) currency() string {
	return "USD"
}

// Franc
type Franc struct {
	amountValue int
}

func (f Franc) amount() int {
	return f.amountValue
}

func (f Franc) times(multiplier int) Money {
	return MakeFranc(f.amountValue * multiplier)
}

func (f Franc) currency() string {
	return "CHF"
}
