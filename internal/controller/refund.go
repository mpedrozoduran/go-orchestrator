package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence"
	db "github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
)

type RefundController struct {
	refundRepository     persistence.Repository[db.Refund]
	auditTrailRepository persistence.Repository[db.AuditTrail]
	coreBankClient       integrations.CoreBankClient
}

func NewRefundController(refundRepo persistence.Repository[db.Refund], auditTrailRepo persistence.Repository[db.AuditTrail], client integrations.CoreBankClient) RefundController {
	return RefundController{
		refundRepository:     refundRepo,
		auditTrailRepository: auditTrailRepo,
		coreBankClient:       client,
	}
}

func (r RefundController) ProcessRefund(request api.RefundRequest) (api.RefundResponse, error) {
	trxReq := integrations.RefundRequest{
		PaymentId: request.PaymentId,
		TrxType:   integrations.TrxTypeRefund,
	}

	refund, err := r.coreBankClient.RefundRequest(trxReq)
	if err != nil {
		return api.RefundResponse{}, nil
	}

	dbRefund := db.Refund{
		KeyId:     refund.ID,
		PaymentId: request.PaymentId,
		Status:    refund.Status,
		Details:   refund.Reason,
		Date:      refund.Date,
	}

	if err := r.refundRepository.Store(dbRefund); err != nil {
		return api.RefundResponse{}, nil
	}

	apiResponse := api.RefundResponse{
		ID:     refund.ID,
		Code:   refund.Code,
		Status: refund.Status,
		Reason: refund.Reason,
		Date:   refund.Date,
	}

	reqData, err := json.Marshal(request)
	if err != nil {
		return api.RefundResponse{}, nil
	}

	resData, err := json.Marshal(apiResponse)
	if err != nil {
		return api.RefundResponse{}, nil
	}

	trail := db.AuditTrail{
		KeyId:    uuid.Must(uuid.NewUUID()).String(),
		Request:  string(reqData),
		Response: string(resData),
	}

	if err = r.auditTrailRepository.Store(trail); err != nil {
		return api.RefundResponse{}, nil
	}

	return apiResponse, nil
}
