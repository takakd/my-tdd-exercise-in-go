package tddexercise

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{}
}

func (d *Dollar)times(multiplier int) {

}