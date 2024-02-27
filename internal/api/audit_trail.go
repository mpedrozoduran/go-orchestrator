package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"net/http"
)

type AuditTrailAPI struct {
	auditTrailController controller.AuditTrailController
}

func NewAuditTrailAPI(controller controller.AuditTrailController) AuditTrailAPI {
	return AuditTrailAPI{
		auditTrailController: controller,
	}
}

func (a AuditTrailAPI) GetAuditTrails(c *gin.Context) {
	res, err := a.auditTrailController.GetAuditTrails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
