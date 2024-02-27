package util

import (
	db "github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
	"time"
)

type MockRepositoryAuditTrail struct {
}

func (m MockRepositoryAuditTrail) Store(data db.AuditTrail) error {
	return nil
}

func (m MockRepositoryAuditTrail) Get(id string) (*db.AuditTrail, error) {
	return nil, nil
}

func (m MockRepositoryAuditTrail) GetAll() ([]db.AuditTrail, error) {
	return []db.AuditTrail{
		{KeyId: "123", Request: "dummy", Response: "dummy", Created: time.Now().Format(DateFormat1)},
	}, nil
}

type MockRepositoryPayment struct {
}

func (m MockRepositoryPayment) Store(data db.Payment) error {
	return nil
}

func (m MockRepositoryPayment) Get(id string) (*db.Payment, error) {
	return &db.Payment{
		KeyId:    "123",
		TrxId:    "123",
		Customer: db.Customer{},
		Merchant: db.Merchant{},
		Status:   "APPROVED",
		Details:  "",
		Date:     "",
	}, nil
}

func (m MockRepositoryPayment) GetAll() ([]db.Payment, error) {
	return nil, nil
}

type MockRepositoryRefund struct {
}

func (m MockRepositoryRefund) Store(data db.Refund) error {
	return nil
}

func (m MockRepositoryRefund) Get(id string) (*db.Refund, error) {
	return nil, nil
}

func (m MockRepositoryRefund) GetAll() ([]db.Refund, error) {
	return nil, nil
}
