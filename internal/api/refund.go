package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"net/http"
)

type RefundsAPI struct {
	RefundController controller.RefundController
}

func NewRefundsAPI(controller controller.RefundController) RefundsAPI {
	return RefundsAPI{
		RefundController: controller,
	}
}

func (r RefundsAPI) ProcessRefund(c *gin.Context) {
	var refund entities.RefundRequest
	if err := c.BindJSON(&refund); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.RefundController.ProcessRefund(refund)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, res)
}
