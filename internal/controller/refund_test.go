package controller

import (
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	appConfig "github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentController_ProcessRefund(t *testing.T) {
	srv := util.TestServer(t)
	defer srv.Close()
	refundRepo := util.MockRepositoryRefund{}
	auditTrailRepo := util.MockRepositoryAuditTrail{}
	cfg := appConfig.Config{
		Server:   appConfig.ServerConfig{Port: 9000},
		Database: appConfig.DatabaseConfig{Url: "mock"},
		Bank:     appConfig.BankConfig{Url: srv.URL},
		Auth: appConfig.AuthConfig{
			Username:  "mock",
			Password:  "mock",
			SecretKey: "mock",
		},
	}
	core := integrations.NewCoreBankClient(cfg)
	controller := NewRefundController(refundRepo, auditTrailRepo, core)
	request := api.RefundRequest{PaymentId: "123"}
	res, err := controller.ProcessRefund(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
}
