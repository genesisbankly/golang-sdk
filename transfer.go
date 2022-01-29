package genesis

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type Transfer struct {
	ID                 int32           `json:"id"`
	Amount             decimal.Decimal `json:"amount"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	Deleted            bool            `json:"-"`
	CompanyID          int32           `json:"company_id"`
	Type               string          `json:"type"`
	Status             string          `json:"status"`
	Internal           bool            `json:"internal"`
	IntegrationID      string          `json:"integration_id"`
	ErrorMessage       string          `json:"error_message"`
	BankNumber         string          `json:"bank_number"`
	BankIspb           string          `json:"bank_ispb"`
	BankName           string          `json:"bank_name"`
	Document           string          `json:"document"`
	Name               string          `json:"name"`
	AccountID          int32           `json:"account_id"`
	AccountType        string          `json:"account_type"`
	Account            string          `json:"account"`
	Branch             string          `json:"branch"`
	Key                string          `json:"key"`
	QrDinamic          string          `json:"qr_dinamic"`
	DigitCode          string          `json:"digit_code"`
	CorrelationID      string          `json:"-"`
	PixKey             string          `json:"pix_key"`
	PixType            string          `json:"pix_type"`
	Protocol           string          `json:"protocol"`
	EndToEndId         string          `json:"EndToEndId"`
	InitializationType string          `json:"InitializationType"`
}

type CreateTransfer struct {
	Type    string `json:"type"`
	PixKey  string `json:"pix_key"`
	PixType string `json:"pix_type"`
	Amount  int    `json:"amount"`
}

type TransferPagesQuery struct {
	PerPage int              `json:"per_page"`
	Page    int              `json:"page"`
	Query   *TransferQuery   `json:"query"`
	OrderBy *TransferOrderBy `json:"order_by"`
}

type TransferQuery struct {
	Status     string `json:"status"`
	To         int64  `json:"to"`
	From       int64  `json:"from"`
	Type       string `json:"type"`
	CompanyID  int32  `json:"company_id"`
	AccountID  int32  `json:"account_id"`
	Internal   bool   `json:"internal"`
	SupplierID int32  `json:"supplier_id"`
}
type TransferOrderBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type TransferPagesResponse struct {
	Data  []*Transfer `json:"data"`
	Page  int         `json:"page"`
	Pages int         `json:"pages"`
	Total int         `json:"total"`
}

// transfer is a structure manager all about transfer
type TransferClient struct {
	client *Genesis
}

//transfer - Instance de transfer
func (g *Genesis) Trasnfer() *TransferClient {
	return &TransferClient{client: g}
}

//Create - create a new wallet
func (d *TransferClient) Create(req CreateTransfer) (*Transfer, *Error, error) {
	data, _ := json.Marshal(req)
	var response *Transfer
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/transfer", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (d *TransferClient) List(req TransferPagesQuery) (*TransferPagesResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *TransferPagesResponse
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/transfers", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
