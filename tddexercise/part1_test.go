package tddexercise

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)

	if 10 != five.amount {
		t.Errorf("expected=%d, actual=%d", five.amount, 10)
	}
}

