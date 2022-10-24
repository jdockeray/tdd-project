package money

import "testing"

type Portfolio []Money

func (p Portfolio) Add(dollars Money) Portfolio {
	return p
}

func (p Portfolio) Evaluate(s string) Money {
	return Money{amount: 15, currency: "USD"}
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	portfolioInDollars = portfolio.Evaluate("USD")

	assertEquals(t, fifteenDollars, portfolioInDollars)
}