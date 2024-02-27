package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	appConfig "github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/integrations"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRefundsAPI_ProcessRefund(t *testing.T) {
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
	controller := controller.NewRefundController(refundRepo, auditTrailRepo, core)
	refundsAPI := NewRefundsAPI(controller)
	router := gin.Default()
	v1 := router.Group("/v1/payments")
	{
		v1.POST("/refund", refundsAPI.ProcessRefund)
	}

	refundReq := entities.RefundRequest{PaymentId: "123"}
	jsonData, err := json.Marshal(refundReq)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/payments/refund", bytes.NewReader(jsonData))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
}
