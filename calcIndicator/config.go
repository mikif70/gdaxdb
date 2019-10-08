package main

import (
	"flag"
	"fmt"
)

type Parameters struct {
	Debug          bool
	candleFilename string
	IndList        string
	indicators     map[string]Indicators
	logFilename    string
}

type Indicators interface {
	InitWindow(window int)
	Init(price float64)
	Calc(price float64) float64
}

var (
	opts = &Parameters{
		indicators: make(map[string]Indicators),
	}
)

func init() {
	flag.StringVar(&opts.logFilename, "l", opts.logFilename, "log filename")
	flag.StringVar(&opts.candleFilename, "f", "", "candle filename")
	flag.StringVar(&opts.IndList, "i", "", "indicators list")
	flag.BoolVar(&opts.Debug, "D", false, "Debug")

	opts.indicators["macd"] = &Macd{}
	opts.indicators["ema"] = &Ema{}
	opts.indicators["sma"] = &Sma{}

	fmt.Printf("opts: %+v\n", opts)
}
