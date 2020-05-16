package tddexercise

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)

	if NewDollar(10).Money != five.times(2) {
		t.Errorf("expected=%d, actual=%d", five.times(2).amount, 10)
	}

	if NewDollar(15).Money != five.times(3) {
		t.Errorf("expected=%d, actual=%d", five.times(3).amount, 15)
	}
}

func Test_Equality(t *testing.T) {
	if NewDollar(5) != NewDollar(5) {
		t.Error("not equal.")
	}
	if NewDollar(5) == NewDollar(6) {
		t.Error("not equal.")
	}
	if NewFranc(5) != NewFranc(5) {
		t.Error("not equal.")
	}
	if NewFranc(5) == NewFranc(6) {
		t.Error("not equal.")
	}

	// compile error: comparing dollar and franc is ok.
	// if NewDollar(5)== NewFranc(5)
	if NewDollar(5).Money == NewFranc(5).Money {
		t.Error("shoud be not equal.")
	}
}

func Test_FrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	if NewFranc(10).Money != five.times(2) {
		t.Errorf("expected=%d, actual=%d", five.times(2).amount, 10)
	}

	if NewFranc(15).Money != five.times(3) {
		t.Errorf("expected=%d, actual=%d", five.times(3).amount, 15)
	}
}

