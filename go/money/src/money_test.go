package main

import "testing"

type Money struct {
	amount   float64
	currency string
}

func (d Money) Times(multiplier int) Money {
	return Money{
		amount:   d.amount * float64(multiplier),
		currency: d.currency,
	}
}

func (d Money) Divide(divisor int) Money {
	return Money{
		amount:   d.amount / float64(divisor),
		currency: d.currency,
	}
}

func assertEquals(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}

	actualMoneyBeforeMultiplication := fiver.Times(2)
	expectedMoneyAfterMultiplication := Money{
		amount:   10,
		currency: "USD",
	}
	assertEquals(t, actualMoneyBeforeMultiplication, expectedMoneyAfterMultiplication)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{
		amount:   10,
		currency: "EUR",
	}

	actualMoneyAfterMultiplication := tenEuros.Times(2)
	expectedMoneyAfterMultiplication := Money{
		amount:   20,
		currency: "EUR",
	}

	assertEquals(t, actualMoneyAfterMultiplication, expectedMoneyAfterMultiplication)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}

	assertEquals(t, actualMoneyAfterDivision, expectedMoneyAfterDivision)
}
