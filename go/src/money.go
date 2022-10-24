package src

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
