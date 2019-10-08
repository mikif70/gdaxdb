package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	var indA []string

	flag.Parse()

	fmt.Printf("indlist: %+v\n", opts.IndList)

	if opts.IndList != "" {
		indA = strings.Split(opts.IndList, ",")
		fmt.Printf("ind: %+v\n", indA)
	}

	fmt.Printf("%+v\n", opts)

	fmt.Printf("%.2f\n", opts.indicators["macd"].Calc(100.00))
	fmt.Printf("%.2f\n", opts.indicators["ema"].Calc(100.00))
	fmt.Printf("%.2f\n", opts.indicators["sma"].Calc(100.00))
}
