package payments

type PaymentMethod interface {
	Pay (amount float64) string
}

