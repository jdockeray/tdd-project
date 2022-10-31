package _go

import (
	s "money/stocks"
	"testing"
)

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioInDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	portfolioInDollars = portfolio.Evaluate("USD")

	assertEquals(t, fifteenDollars, portfolioInDollars)
}
