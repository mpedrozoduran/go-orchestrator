package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuditTrailAPI_GetAuditTrails(t *testing.T) {
	srv := util.TestServer(t)
	defer srv.Close()
	auditTrailRepo := util.MockRepositoryAuditTrail{}
	controller := controller.NewAuditTrailController(auditTrailRepo)
	auditTrailsAPI := NewAuditTrailAPI(controller)
	router := gin.Default()
	v1 := router.Group("/v1/payments")
	{
		v1.GET("/history", auditTrailsAPI.GetAuditTrails)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/payments/history", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
