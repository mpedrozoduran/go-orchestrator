package controller

import (
	"encoding/json"
	"fmt"
	appConfig "github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(t *testing.T) *httptest.Server {
	fmt.Println("mocking server")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := struct {
			TrxID  string `json:"trx_id"`
			Status string `json:"status"`
			Code   int    `json:"code"`
			Date   string `json:"date"`
		}{
			TrxID:  "abc123",
			Status: "APPROVED",
			Code:   10001,
			Date:   "",
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			assert.NoError(t, err)
			return
		}
		_, err = w.Write(jsonData)
		if err != nil {
			assert.NoError(t, err)
			return
		}
	}))
	return server
}

func TestPaymentController_ProcessPayment(t *testing.T) {
	srv := testServer(t)
	defer srv.Close()
	paymentRepo := util.MockRepositoryPayment{}
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
	controller := NewPaymentController(paymentRepo, auditTrailRepo, core)
	request := util.BuildMockPayment()
	res, err := controller.ProcessPayment(request)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
}

func TestPaymentController_GetPayment(t *testing.T) {
	srv := testServer(t)
	defer srv.Close()
	paymentRepo := util.MockRepositoryPayment{}
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
	controller := NewPaymentController(paymentRepo, auditTrailRepo, core)
	res, err := controller.GetPayment("abc123")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
