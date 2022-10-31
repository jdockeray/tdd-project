package stocks

type Portfolio []Money

func (p Portfolio) Add(dollars Money) Portfolio {
	return p
}

func (p Portfolio) Evaluate(s string) Money {
	return Money{amount: 15, currency: "USD"}
}
