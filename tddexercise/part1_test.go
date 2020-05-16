package tddexercise

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)

	product := five.times(2)
	if 10 != product.amount {
		t.Errorf("expected=%d, actual=%d", product.amount, 10)
	}

	product = five.times(3)
	if 15 != product.amount {
		t.Errorf("expected=%d, actual=%d", product.amount, 15)
	}
}

func Test_Equality(t *testing.T) {
	if NewDollar(5) != NewDollar(5) {
		t.Error("not equal.")
	}
	if NewDollar(5) == NewDollar(6) {
		t.Error("not equal.")
	}
}
