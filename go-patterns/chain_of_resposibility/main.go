package main

import "fmt"

type Handler interface {
	Handle(byte)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h ConcreteHandlerA) Handle(msg byte) {
	if msg == 'A' {
		fmt.Println("The request was handled by ConcreteHandlerA")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h ConcreteHandlerB) Handle(msg byte) {
	if msg == 'B' {
		fmt.Println("The request was handled by ConcreteHandlerB")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

type ConcreteHandlerC struct {
	next Handler
}

func (h ConcreteHandlerC) Handle(msg byte) {
	if msg == 'C' {
		fmt.Println("The request was handled by ConcreteHandlerC")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

func main() {
	handler := ConcreteHandlerA{
		next: ConcreteHandlerB{
			next: ConcreteHandlerC{},
		},
	}

	handler.Handle('C')
}
