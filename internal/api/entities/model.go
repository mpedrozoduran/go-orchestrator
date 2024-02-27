package entities

type Payment struct {
	Customer Customer `json:"customer"`
	Merchant Merchant `json:"merchant"`
	Amount   float64  `json:"amount"`
}

type PaymentRequest struct {
	Payment
}

type PaymentResponse struct {
	ID       string   `json:"id"`
	TrxID    string   `json:"trx_id"`
	Customer Customer `json:"customer"`
	Merchant Merchant `json:"merchant"`
	Status   string   `json:"status"`
	Details  string   `json:"details"`
	Date     string   `json:"date"`
}

type RefundRequest struct {
	PaymentId string `json:"payment_id"`
}

type RefundResponse struct {
	ID     string `json:"trx_id"`
	Code   int    `json:"code"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Date   string `json:"date"`
}

type Customer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Card     Card   `json:"card"`
}

type Card struct {
	Number          string `json:"number"`
	ExpirationMonth string `json:"exp_month"`
	ExpirationYear  string `json:"exp_year"`
	CVV             string `json:"cvv"`
}

type Merchant struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	AccountNumber string `json:"account_number"`
	BankId        string `json:"bank_id"`
}

type AuditTrail struct {
	ID       string `json:"id"`
	Request  string `json:"request"`
	Response string `json:"response"`
	Created  string `json:"created"`
}
