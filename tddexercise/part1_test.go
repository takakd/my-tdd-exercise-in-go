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

	bank := NewBank()
	reduced := bank.reduce(sum, "USD")
	if MakeDollar(10) != reduced {
		t.Errorf("expected=%d, actual=%d", MakeDollar(10).amount, reduced.amount)
	}
}

func Test_PlusReturnSum(t *testing.T) {
	five := MakeDollar(5)
	result := five.plus(five)
	// @note: interface is pointer.
	sum, _ := result.(*Sum)
	if five != sum.augend {
		t.Errorf("expected=%d, actual=%d", five.amount, sum.augend.amount)
	}
	if five != sum.addend {
		t.Errorf("expected=%d, actual=%d", five.amount, sum.addend.amount)
	}
}

func Test_ReduceSum(t *testing.T) {
	sum := NewSum(MakeDollar(3), MakeDollar(4))
	bank := NewBank()
	result := bank.reduce(sum, "USD")
	if result != MakeDollar(7) {
		t.Errorf("expected=%d, actual=%d", result.amount, MakeDollar(7).amount)
	}
}

func Test_ReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.reduce(MakeDollar(1), "USD")
	if MakeDollar(1) != result {
		t.Errorf("expected=%d, actual=%d", result.amount, MakeDollar(1).amount)
	}
}

func Test_ReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.reduce(MakeFranc(2), "USD")
	if MakeDollar(1) != result {
		t.Errorf("expected=%d, actual=%d", result.amount, MakeDollar(1).amount)
	}
}

func Test_ArrayEquals(t *testing.T) {
	if [...]string{"abc"} != [...]string{"abc"} {
		t.Error("not equal array")
	}
}
