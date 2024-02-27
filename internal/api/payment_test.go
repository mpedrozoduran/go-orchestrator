package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	appConfig "github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentsAPI_ProcessPayment(t *testing.T) {
	srv := util.TestServer(t)
	defer srv.Close()
	refundRepo := util.MockRepositoryPayment{}
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
	controller := controller.NewPaymentController(refundRepo, auditTrailRepo, core)
	paymentsAPI := NewPaymentsAPI(controller)
	router := gin.Default()
	v1 := router.Group("/v1/payments")
	{
		v1.POST("", paymentsAPI.ProcessPayment)
	}

	paymentReq := util.BuildMockPayment()
	jsonData, err := json.Marshal(paymentReq)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/payments", bytes.NewReader(jsonData))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPaymentsAPI_GetPayment(t *testing.T) {
	srv := util.TestServer(t)
	defer srv.Close()
	refundRepo := util.MockRepositoryPayment{}
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
	controller := controller.NewPaymentController(refundRepo, auditTrailRepo, core)
	paymentsAPI := NewPaymentsAPI(controller)
	router := gin.Default()
	v1 := router.Group("/v1/payments")
	{
		v1.GET("/:id", paymentsAPI.GetPayment)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payments/123", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
