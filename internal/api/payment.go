package api

import (
	"github.com/gin-gonic/gin"
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/controller"
	"net/http"
)

type PaymentsAPI struct {
	paymentController controller.PaymentController
}

func NewPaymentsAPI(controller controller.PaymentController) PaymentsAPI {
	return PaymentsAPI{
		paymentController: controller,
	}
}

func (p PaymentsAPI) ProcessPayment(c *gin.Context) {
	var payment api.Payment
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.paymentController.ProcessPayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (p PaymentsAPI) GetPayment(c *gin.Context) {
	paymentId := c.Param("id")
	if len(paymentId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment id must not be empty"})
		return
	}

	res, err := p.paymentController.GetPayment(paymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res.Status != "" {
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusNotFound, nil)
}
