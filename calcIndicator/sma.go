package main

import (
	"fmt"
	"log"
)

type Sma struct {
	//	interval string
	window      int
	windowClose []float64
	Sma         float64
}

func (s *Sma) Update(price float64) {
	l := len(s.windowClose)
	for i := 1; i < l; i++ {
		s.windowClose[i-1] = s.windowClose[i]
	}
	log.Printf("sma newval: len %d\n", l)
	s.windowClose[l-1] = price

	s.calcSma()
}

func (s *Sma) calcSma() {

	var sma float64
	var length int

	if len(s.windowClose) == 0 {
		fmt.Println("Sma vals empty")
		return
	}

	if len(s.windowClose) < s.window {
		length = len(s.windowClose)
	} else {
		length = s.window
	}

	for i := length; i > 0; i-- {

		sma += s.windowClose[i-1]
	}

	s.Sma = sma / float64(length)
}

func (s *Sma) Calc(price float64) float64 {
	return price * 9.0
}

func (s *Sma) InitWindow(window int) {
	s.window = window
	s.windowClose = make([]float64, window)
}

func (s *Sma) Init(price float64) {

}
