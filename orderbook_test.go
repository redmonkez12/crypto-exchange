package main

import "testing"

func TestOrderBook(t *testing.T) {
	l := NewLimitOrder(10_000)
	buyOrder := NewOrder(true, 5)

	l.AddOrder(buyOrder)
}
