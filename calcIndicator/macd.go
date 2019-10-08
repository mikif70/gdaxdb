package main

//	"fmt"
//	"log"

var (
	_macdWindow = []int{10, 26, 9}
)

type Macd struct {
	window    int
	line      []float64
	signal    float64
	histogram float64
}

// Macd line = 10-day EMA - 26-day EMA
// Macd signal = media ultimi 9 Macd line
// MACD Histogram = Macd line - Macd signal
func NewMacd(candlePeriod int, currency string) *Macd {
	macd := &Macd{
		line:   make([]float64, _macdWindow[2]),
		window: _macdWindow[2],
	}

	//	macd.calcMacd(ema1, ema2)

	return macd
}

func (m *Macd) calcLine(ema1 float64, ema2 float64) {
	l := len(m.line)
	for i := 1; i < l; i++ {
		m.line[i-1] = m.line[i]
	}

	m.line[l-1] = ema1 - ema2
}

func (m *Macd) calcSignal() {
	var signal float64

	for i := m.window; i > 0; i-- {
		signal += m.line[i-1]
	}

	m.signal = signal / float64(m.window)
}

func (m *Macd) calcHistogram() {
	m.histogram = m.line[m.window-1] - m.signal
}

func (m *Macd) Update(ema1 float64, ema2 float64) {
	if ema1 != 0 && ema2 != 0 {
		m.calcLine(ema1, ema2)
	}

	m.calcSignal()

	m.calcHistogram()
}

func (m *Macd) Calc(price float64) float64 {
	return price * 6.0
}

func (m *Macd) InitWindow(window int) {
	m.window = window
	m.line = make([]float64, window)
}

func (m *Macd) Init(price float64) {

}
