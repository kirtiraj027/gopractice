import payment

import "fmt"

type payment struct {
	CreditNumber string
	DebitNumber string
	QR code string



func (c payment) Pay
(amount float64) string {
	msg := fmt.Sprintf(" Payment is of successful", amount, c.payment)
	return msg

}
