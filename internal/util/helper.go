package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupGinRouter() *gin.Engine {
	r := gin.Default()
	r.Group("/v1/payments")
	{
		r.POST("", func(c *gin.Context) {
			c.String(201, "pong")
		})
		r.GET("/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})
		r.GET("/refund/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})
		r.GET("/history", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}

	return r
}

func TestServer(t *testing.T) *httptest.Server {
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

func BuildMockPayment() api.Payment {
	return api.Payment{
		Customer: api.Customer{
			ID:       "123",
			Name:     "dummy",
			Lastname: "dummy",
			Card: api.Card{
				Number:          "123",
				ExpirationMonth: "01",
				ExpirationYear:  "2026",
				CVV:             "000",
			},
		},
		Merchant: api.Merchant{
			ID:            "123",
			Name:          "dummy",
			Lastname:      "dummy",
			AccountNumber: "123",
			BankId:        "1",
		},
		Amount: 10,
	}
}
