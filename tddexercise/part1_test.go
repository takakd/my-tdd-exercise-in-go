package tddexercise

import (
	"testing"
	"reflect"
	"fmt"
)

func Test_Multiplication(t *testing.T) {
	five := MakeDollar(5)

	fmt.Printf("%v, %v", MakeDollar(10), five.times(2))
	if ! reflect.DeepEqual(MakeDollar(10), five.times(2)) {
		t.Errorf("expected=%d, actual=%d", five.times(2).amount, 10)
	}

	if ! reflect.DeepEqual(MakeDollar(15), five.times(3)) {
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
	if MakeFranc(5) != MakeFranc(5) {
		t.Error("not equal.")
	}
	if MakeFranc(5) == MakeFranc(6) {
		t.Error("not equal.")
	}

	// compile error: comparing dollar and franc is ok.
	// if MakeDollar(5)== MakeFranc(5)
	if MakeDollar(5) == MakeFranc(5) {
		t.Error("shoud be not equal.")
	}
}

func Test_FrancMultiplication(t *testing.T) {
	five := MakeFranc(5)

	if MakeFranc(10) != five.times(2) {
		t.Errorf("expected=%d, actual=%d", five.times(2).amount, 10)
	}

	if MakeFranc(15) != five.times(3) {
		t.Errorf("expected=%d, actual=%d", five.times(3).amount, 15)
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
