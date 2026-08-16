package main

import "fmt"

type Payment interface {
	Pay(amount int)
}

type CreditCard struct {}
func (c CreditCard)Pay(amount int){
	fmt.Println("paid amount:", amount, "by credit card")
}

type Cash struct {}
func (cash Cash)Pay(amount int){
	fmt.Println("paid amount:", amount, "by cash")
}

type PaymentGateway struct {
	method Payment
}

func NewPaymentGateway(method Payment) *PaymentGateway {
	return &PaymentGateway{
		method: method,
	}
}

func (g *PaymentGateway) SetPaymentMethod (method Payment){
	g.method = method
}

func (g *PaymentGateway) Checkout(amount int){
	g.method.Pay(amount)
}

func main(){
	gateway := NewPaymentGateway(nil)

	gateway.SetPaymentMethod(CreditCard{})
	gateway.Checkout(100)

	gateway.SetPaymentMethod(Cash{})
	gateway.Checkout(200)
	
}