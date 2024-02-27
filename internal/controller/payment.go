package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence"
	db "github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"time"
)

type PaymentController struct {
	paymentsRepository   persistence.Repository[db.Payment]
	auditTrailRepository persistence.Repository[db.AuditTrail]
	coreBankClient       integrations.CoreBankClient
}

func NewPaymentController(paymentsRepo persistence.Repository[db.Payment], auditTrailRepo persistence.Repository[db.AuditTrail], client integrations.CoreBankClient) PaymentController {
	return PaymentController{
		paymentsRepository:   paymentsRepo,
		auditTrailRepository: auditTrailRepo,
		coreBankClient:       client,
	}
}

func (p PaymentController) ProcessPayment(request api.Payment) (api.PaymentResponse, error) {
	trxReq := integrations.TransactionRequest{
		Transaction: integrations.Transaction{
			Customer: integrations.Customer{
				ID:       request.Customer.ID,
				Name:     request.Customer.Name,
				Lastname: request.Customer.Lastname,
				Card: integrations.Card{
					Number:          request.Customer.Card.Number,
					ExpirationMonth: request.Customer.Card.ExpirationMonth,
					ExpirationYear:  request.Customer.Card.ExpirationYear,
					CVV:             request.Customer.Card.CVV,
				},
			},
			Merchant: integrations.Merchant{
				ID:            request.Merchant.ID,
				Name:          request.Merchant.Name,
				Lastname:      request.Merchant.Lastname,
				AccountNumber: request.Merchant.AccountNumber,
				BankId:        request.Merchant.BankId,
			},
			Amount:  request.Amount,
			TrxType: integrations.TrxTypePayment,
		},
	}

	payment, err := p.coreBankClient.PaymentRequest(trxReq)
	if err != nil {
		return api.PaymentResponse{}, err
	}

	dbPayment := db.Payment{
		KeyId: uuid.Must(uuid.NewUUID()).String(),
		TrxId: payment.ID,
		Customer: db.Customer{
			Id:       request.Customer.ID,
			Name:     request.Customer.Name,
			Lastname: request.Customer.Lastname,
			Card: db.Card{
				Number:          request.Customer.Card.Number,
				ExpirationMonth: request.Customer.Card.ExpirationMonth,
				ExpirationYear:  request.Customer.Card.ExpirationYear,
				CVV:             request.Customer.Card.CVV,
			},
		},
		Merchant: db.Merchant{
			Id:            request.Merchant.ID,
			Name:          request.Merchant.Name,
			Lastname:      request.Merchant.Lastname,
			AccountNumber: request.Merchant.AccountNumber,
			BankId:        request.Merchant.BankId,
		},
		Status:  payment.Status,
		Details: payment.Reason,
		Date:    payment.Date,
	}

	if err = p.paymentsRepository.Store(dbPayment); err != nil {
		return api.PaymentResponse{}, err
	}

	apiPayment := api.PaymentResponse{
		ID:    dbPayment.KeyId,
		TrxID: payment.ID,
		Customer: api.Customer{
			ID:       request.Customer.ID,
			Name:     request.Customer.Name,
			Lastname: request.Customer.Lastname,
			Card: api.Card{
				Number:          request.Customer.Card.Number,
				ExpirationMonth: request.Customer.Card.ExpirationMonth,
				ExpirationYear:  request.Customer.Card.ExpirationYear,
				CVV:             request.Customer.Card.CVV,
			},
		},
		Merchant: api.Merchant{
			ID:            request.Merchant.ID,
			Name:          request.Merchant.Name,
			Lastname:      request.Merchant.Lastname,
			AccountNumber: request.Merchant.AccountNumber,
			BankId:        request.Merchant.BankId,
		},
		Status:  payment.Status,
		Details: payment.Reason,
		Date:    payment.Date,
	}

	reqData, err := json.Marshal(request)
	if err != nil {
		return api.PaymentResponse{}, err
	}

	resData, err := json.Marshal(apiPayment)
	if err != nil {
		return api.PaymentResponse{}, err
	}

	trail := db.AuditTrail{
		KeyId:    uuid.Must(uuid.NewUUID()).String(),
		Request:  string(reqData),
		Response: string(resData),
		Created:  time.Now().Format(util.DateFormat1),
	}

	if err = p.auditTrailRepository.Store(trail); err != nil {
		return api.PaymentResponse{}, err
	}

	return apiPayment, nil
}

func (p PaymentController) GetPayment(id string) (api.PaymentResponse, error) {
	res, err := p.paymentsRepository.Get(id)
	if err != nil {
		return api.PaymentResponse{}, err
	}
	if res == nil {
		return api.PaymentResponse{}, nil
	}
	apiPayment := api.PaymentResponse{
		ID:    res.KeyId,
		TrxID: res.TrxId,
		Customer: api.Customer{
			ID:       res.Customer.Id,
			Name:     res.Customer.Name,
			Lastname: res.Customer.Lastname,
			Card: api.Card{
				Number:          res.Customer.Card.Number,
				ExpirationMonth: res.Customer.Card.ExpirationMonth,
				ExpirationYear:  res.Customer.Card.ExpirationYear,
				CVV:             res.Customer.Card.CVV,
			},
		},
		Merchant: api.Merchant{
			ID:            res.Merchant.Id,
			Name:          res.Merchant.Name,
			Lastname:      res.Merchant.Lastname,
			AccountNumber: res.Merchant.AccountNumber,
			BankId:        res.Merchant.BankId,
		},
		Status:  res.Status,
		Details: res.Details,
		Date:    res.Date,
	}
	return apiPayment, nil
}
