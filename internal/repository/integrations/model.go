package integrations

type TrxType int

const (
	TrxTypePayment TrxType = iota
	TrxTypeRefund
)

type Transaction struct {
	Customer `json:"customer"`
	Merchant `json:"merchant"`
	Amount   float64 `json:"amount"`
	TrxType  TrxType `json:"trx_type"`
}

type TransactionRequest struct {
	Transaction
}

type TransactionResponse struct {
	ID     string `json:"trx_id"`
	Code   int    `json:"code"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Date   string `json:"date"`
}

type RefundRequest struct {
	PaymentId string  `json:"payment_id"`
	TrxType   TrxType `json:"trx_type"`
}

type RefundResponse struct {
	ID     string `json:"trx_id"`
	Code   string `json:"code"`
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
	ID  string `json:"id"`
	Log string `json:"log"`
}
