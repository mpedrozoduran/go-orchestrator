package controller

import (
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentController_GetAuditTrails(t *testing.T) {
	srv := util.TestServer(t)
	defer srv.Close()
	auditTrailRepo := util.MockRepositoryAuditTrail{}
	controller := NewAuditTrailController(auditTrailRepo)
	res, err := controller.GetAuditTrails()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
