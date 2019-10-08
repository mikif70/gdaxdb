// ema
package main

type Ema struct {
	window     int
	multiplier float64
	Ema        float64
}

// [2 ÷ (selected time period + 1)]
func (e *Ema) calcWeight() {
	//	w = 2.0 / (period + 1.0)
	e.multiplier = (float64(2) / float64(e.window+1))
}

// [Closing price - EMA(prev)] x multiplier + EMA(prev)
// t0 EMA(prev) = SMA
func (e *Ema) Update(price float64) {
	e.Ema = (price-e.Ema)*e.multiplier + e.Ema
}

func (e *Ema) Calc(price float64) float64 {
	return price * 5.0
}

func (e *Ema) InitWindow(window int) {
	e.window = window
}

func (e *Ema) Init(sma float64) {
	e.Ema = sma

	e.calcWeight()
}
