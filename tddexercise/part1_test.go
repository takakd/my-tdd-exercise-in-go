package tddexercise

import (
	"testing"
	"reflect"
)

func Test_Multiplication(t *testing.T) {
	five := MakeDollar(5)

	if MakeDollar(10) != five.times(2) {
		t.Errorf("expected=%d, actual=%d", five.times(2).amount, 10)
	}

	if MakeDollar(15) != five.times(3) {
		t.Errorf("expected=%d, actual=%d", five.times(3).amount, 15)
	}
}

func Test_Equality(t *testing.T) {
	if MakeDollar(5) != MakeDollar(5) {
		t.Error("not equal.")
	}
	if MakeDollar(5) == MakeDollar(6) {
		t.Error("not equal.")
	}
	if reflect.DeepEqual(MakeDollar(5), MakeFranc(5)) {
		t.Error("shoud be not equal.")
	}
}

func Test_Currency(t *testing.T) {
	if "USD" != MakeDollar(1).currency {
		t.Errorf("expected=%s, actual=%s", "USD", MakeDollar(1).currency)
	}
	if "CHF" != MakeFranc(1).currency {
		t.Errorf("expected=%s, actual=%s", "CHF", MakeFranc(1).currency)
	}
}

func Test_SimpleAddition(t *testing.T) {
	sum := MakeDollar(5).plus(MakeDollar(5))
	if sum != MakeDollar(10) {
		t.Errorf("expected=%d, actual=%d", MakeDollar(10).amount, sum.amount)
	}
}

//
func Test_SimpleAddition(t *testing.T) {
	bank := NewBank()
	reduced := bank.reduce(sum, "USD")
	if MakeDollar(10) != reduced {
		t.Errorf("expected=%d, actual=%d", MakeDollar(10).amount, reduced)
	}
}
