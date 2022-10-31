package _go

import (
	s "money/stocks"
	"testing"
)

func assertEquals(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	fiver := s.NewMoney(5, "USD")

	actualMoneyBeforeMultiplication := fiver.Times(2)
	expectedMoneyAfterMultiplication := s.NewMoney(10, "USD")

	assertEquals(t, actualMoneyBeforeMultiplication, expectedMoneyAfterMultiplication)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")

	actualMoneyAfterMultiplication := tenEuros.Times(2)
	expectedMoneyAfterMultiplication := s.NewMoney(20, "EUR")

	assertEquals(t, actualMoneyAfterMultiplication, expectedMoneyAfterMultiplication)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.NewMoney(1000.5, "KRW")

	assertEquals(t, actualMoneyAfterDivision, expectedMoneyAfterDivision)
}
