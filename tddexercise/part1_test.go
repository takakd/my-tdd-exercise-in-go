package tddexercise

import (
	"testing"
	"reflect"
)

func Test_Multiplication(t *testing.T) {
	five := MakeDollar(5)

	if MakeDollar(10) != five.times(2) {
		t.Errorf("expected=%d, actual=%d", five.times(2), 10)
	}

	if MakeDollar(15) != five.times(3) {
		t.Errorf("expected=%d, actual=%d", five.times(3), 15)
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
		t.Errorf("expected=%d, actual=%d", five.amount, sum.augend)
	}
	if five != sum.addend {
		t.Errorf("expected=%d, actual=%d", five.amount, sum.addend)
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

func Test_MapEqual(t *testing.T) {
	l := map[struct {
		from string
		to   string
	}]int{{"CHF", "USD"}: 2}
	r := map[struct {
		from string
		to   string
	}]int{{"CHF", "USD"}: 2}

	if ! reflect.DeepEqual(l, r) {
		t.Error("not equal array")
	}
}

func Test_IdentityRate(t *testing.T) {
	bank := NewBank()
	if 1 != bank.rate("USD", "USD") {
		t.Errorf("expected=%d, actual=%d", 1, bank.rate("USD", "USD"))
	}
}

func Test_MixedAddition(t *testing.T) {
	fiveBucks := Expression(MakeDollar(5))
	tenFrancs := Expression(MakeFranc(10))
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.reduce(fiveBucks.plus(tenFrancs), "USD")
	if result != MakeDollar(10) {
		t.Errorf("expected=%d, actual=%d", result.amount, MakeDollar(10).amount)
	}
}

func Test_SumPlusMoney(t *testing.T) {
	fiveBucks := MakeDollar(5)
	tenFrancs := MakeFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	// this test for "Sum", so we can not write "fiveBucks.plus(tenFrancs)"
	sum := NewSum(fiveBucks, tenFrancs).plus(fiveBucks)
	result := bank.reduce(sum, "USD")
	if MakeDollar(15) != result {
		t.Errorf("expected=%d, actual=%d", MakeDollar(15).amount, result.amount)
	}
}

func Test_SumTimes(t *testing.T) {
	fiveBucks := MakeDollar(5)
	tenFrancs := MakeFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).times(2)
	result := bank.reduce(sum, "USD")
	if MakeDollar(20) != result {
		t.Errorf("expected=%d, actual=%d", MakeDollar(20).amount, result.amount)
	}
}

// can not think of a good way
//func Test_PlusSameCurrencyReturnMoney(t *testing.T) {
//	sum := MakeDollar(5).plus(MakeDollar(5))
//	_, ok := sum.(Money)
//	if ! ok {
//		t.Error("type error. type should be Money.")
//	}
//}
