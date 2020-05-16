package tddexercise

// Money
type IMoney interface {
	amount() int
	times(multiplier int) IMoney
	currency() string
}

func MakeDollar(amount int) IMoney {
	return Dollar{Money{amountValue: amount, currencyValue: "USD"}}
}

func MakeFranc(amount int) IMoney {
	return Franc{Money{amountValue: amount, currencyValue: "CHF"}}
}

// common impl
type Money struct {
	amountValue   int
	currencyValue string
}

func (m Money) amount() int {
	return m.amountValue
}

func (m Money) currency() string {
	return m.currencyValue
}

func (m Money) times(multiplier int) IMoney {
	return m
}

// Dollar
type Dollar struct {
	Money
}

func (d Dollar) amount() int {
	return d.amountValue
}

func (d Dollar) times(multiplier int) IMoney {
	return MakeDollar(d.amountValue * multiplier)
}

// Franc
type Franc struct {
	Money
}

func (f Franc) amount() int {
	return f.amountValue
}

func (f Franc) times(multiplier int) IMoney {
	return MakeFranc(f.amountValue * multiplier)
}
