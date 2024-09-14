package old_examples

import (
	"fmt"
)

type Order struct {
	ID   int
	Type string
}

type Offer struct {
	ID int
}

type Collectable interface {
	Order | Offer | string
}

type Collection[T Collectable] struct {
	collection []T
	index      int
}

func (c *Collection[T]) addItem(item T) {
	c.collection = append(c.collection, item)
}

func (c *Collection[T]) hasNext() bool {
	if c.index < len(c.collection) {
		return true
	}
	return false
}

func (c *Collection[T]) getNext() *T {
	if c.hasNext() {
		result := c.collection[c.index]
		c.index++
		return &result
	}
	return nil
}

func main_col() {
	myOrderCollection := &Collection[Order]{}
	order := Order{1, "order"}
	order2 := Order{2, "order"}
	myOrderCollection.addItem(order)
	myOrderCollection.addItem(order2)
	for myOrderCollection.hasNext() {
		item := myOrderCollection.getNext()
		fmt.Println(*item)
	}

	myOfferCollection := &Collection[Offer]{}
	offer1 := Offer{3}
	offer2 := Offer{4}
	myOfferCollection.addItem(offer1)
	myOfferCollection.addItem(offer2)
	for myOfferCollection.hasNext() {
		item := myOfferCollection.getNext()
		fmt.Println(*item)
	}

	myStringCollection := &Collection[string]{}
	myStringCollection.addItem("string1")
	myStringCollection.addItem("string2")
	for myStringCollection.hasNext() {
		item := myStringCollection.getNext()
		fmt.Println(*item)
	}
}
