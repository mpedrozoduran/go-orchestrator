package entities

type Payment struct {
	KeyId    string   `dynamo:"KeyId"`
	TrxId    string   `dynamo:"TrxId"`
	Customer Customer `dynamo:"Customer"`
	Merchant Merchant `dynamo:"Merchant"`
	Status   string   `dynamo:"Status"`
	Details  string   `dynamo:"Details"`
	Date     string   `dynamo:"Date"`
}

type Refund struct {
	KeyId     string `dynamo:"KeyId"`
	PaymentId string `dynamo:"PaymentId"`
	Status    string `dynamo:"Status"`
	Details   string `dynamo:"Details"`
	Date      string `dynamo:"Date"`
}

type Customer struct {
	Id       string `dynamo:"Id"`
	Name     string `dynamo:"Name"`
	Lastname string `dynamo:"Lastname"`
	Card     Card   `dynamo:"Card"`
}

type Card struct {
	Number          string `dynamo:"Number"`
	ExpirationMonth string `dynamo:"ExpMonth"`
	ExpirationYear  string `dynamo:"ExpYear"`
	CVV             string `dynamo:"Cvv"`
}

type Merchant struct {
	Id            string `dynamo:"Id"`
	Name          string `dynamo:"Name"`
	Lastname      string `dynamo:"Lastname"`
	AccountNumber string `dynamo:"AccountNumber"`
	BankId        string `dynamo:"BankId"`
}

type AuditTrail struct {
	KeyId    string `dynamo:"KeyId"`
	Request  string `dynamo:"Request"`
	Response string `dynamo:"Response"`
	Created  string `dynamo:"Created"`
}
