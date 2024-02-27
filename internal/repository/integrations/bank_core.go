package integrations

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/mpedrozoduran/go-orchestrator/internal/config"
	"github.com/mpedrozoduran/go-orchestrator/internal/util"
	"io"
	"net/http"
)

type CoreBankClient struct {
	appConfig config.Config
}

func NewCoreBankClient(appConfig config.Config) CoreBankClient {
	return CoreBankClient{appConfig: appConfig}
}

func (c CoreBankClient) request(url string, data []byte) (TransactionResponse, error) {
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return TransactionResponse{}, err
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return TransactionResponse{}, err
	}
	defer resp.Body.Close()
	var trxResp TransactionResponse
	if err := json.Unmarshal(out, &trxResp); err != nil {
		return TransactionResponse{}, err
	}
	return trxResp, nil
}

func (c CoreBankClient) PaymentRequest(request TransactionRequest) (TransactionResponse, error) {
	url := fmt.Sprintf("%s%s", c.appConfig.Bank.Url, util.MockBankPaymentSuccessEndpoint)
	data, err := json.Marshal(request)
	if err != nil {
		return TransactionResponse{}, nil
	}
	return c.request(url, data)
}

func (c CoreBankClient) RefundRequest(request RefundRequest) (TransactionResponse, error) {
	url := fmt.Sprintf("%s%s", c.appConfig.Bank.Url, util.MockBankRefundSuccessEndpoint)
	data, err := json.Marshal(request)
	if err != nil {
		return TransactionResponse{}, nil
	}
	return c.request(url, data)
}
