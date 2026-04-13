package main

import "fmt"



type Payment interface {
	Pay() string
}

type CreditCard struct{}

func (c CreditCard) Pay() string {
	return "Paid using Credit Card"
}
type DebitCard struct{}

func (d DebitCard) Pay() string {
	return "Paid using Debit Card"
}
type UPI struct{}

func (u UPI) Pay() string {
	return "Paid using UPI"
}
type QRScanner struct{}

func (q QRScanner) Pay() string {
	return "Paid using QR Scanner"
}

func main() {
	var p Payment

	p = CreditCard{}
	fmt.Println(p.Pay())

	p = DebitCard{}
	fmt.Println(p.Pay())

	p = UPI{}
	fmt.Println(p.Pay())

	p = QRScanner{}
	fmt.Println(p.Pay())
}
