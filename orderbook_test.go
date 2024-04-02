package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimitOrder(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 5)
	buyOrderC := NewOrder(true, 5)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}

func TestOrderBook(t *testing.T) {
	ob := NewOrderBook()

	buyOrderA := NewOrder(true, 10)
	buyOrderB := NewOrder(true, 2000)
	ob.PlaceOrder(18_000, buyOrderA)
	ob.PlaceOrder(19_000, buyOrderB)

	//for i := 0; i < len(ob.Bids); i++ {
	//	fmt.Printf("%+v", ob.Bids[i])
	//}

	fmt.Println(ob.Bids)
}
