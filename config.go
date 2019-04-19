package main

import (
	"flag"
)

// Parameters call parameters
type Parameters struct {
	Product     string
	Filename    string
	Start       string
	Granularity int
	Iteration   int
	Append      bool
	Timestamp   bool
}

var (
	parameters = &Parameters{}
)

func init() {
	flag.StringVar(&parameters.Product, "p", "BTC-EUR", "product")
	flag.StringVar(&parameters.Filename, "f", "", "output filename")
	flag.StringVar(&parameters.Start, "s", "", "start date <2006-01-02T15:04:05Z>")
	flag.IntVar(&parameters.Granularity, "g", 3600, "granularity")
	flag.IntVar(&parameters.Iteration, "i", 1, "iteration")
	flag.BoolVar(&parameters.Append, "a", false, "Append to filename")
	flag.BoolVar(&parameters.Timestamp, "t", false, "Save time as timestamp")
}
