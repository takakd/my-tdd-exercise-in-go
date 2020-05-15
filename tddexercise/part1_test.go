package tddexercise

import "testing"

func testMultiplication(t testing.T) {
	five := NewDollar(5)
	five.times(2)

	if 10 != five.amount {
		t.Errorf(`expected=%s, actual=%s`, five.amount, 10)
	}
}

