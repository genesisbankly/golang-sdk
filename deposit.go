package genesis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Deposit struct {
	ID               int32           `json:"id"`
	Protocol         string          `json:"protocol"`
	Type             string          `json:"type"`
	Amount           decimal.Decimal `gorm:"column:amount" sql:"type:decimal(32,2);" json:"amount"`
	AmountRequest    decimal.Decimal `gorm:"column:amount_request" sql:"type:decimal(32,2);" json:"amount_request"`
	Account          string          `json:"account"`
	Branch           string          `json:"branch"`
	Status           string          `json:"status"`
	IntegrationID    string          `json:"integration_id"`
	LinkRef          string          `json:"link_ref" gorm:"type:text"`
	DocumentRef      string          `json:"document_ref"`
	Source           string          `json:"source"`
	TransactionID    int32           `json:"transaction_id"`
	DocumentSender   string          `json:"document_sender"`
	ExpireDate       time.Time       `json:"expire_date"`
	BarCode          string          `json:"bar_code"`
	DigitableLine    string          `json:"digitable_line"`
	ProcessingTime   int64           `json:"processing_time"`
	SenderBranch     string          `json:"sender_branch"`
	SenderNumber     string          `json:"sender_number"`
	SenderDocument   string          `json:"sender_document"`
	SenderIspbNumber string          `json:"sender_ispb_number"`
	SenderName       string          `json:"sender_name"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
	ProcessedAt      time.Time       `json:"processedAt"`
	BankName         string          `json:"bank_name"`
	EndToEndId       string          `json:"endToEndId"`
	BrcodeKey        string          `json:"brcodekey"`
	BrcodeID         int32           `json:"brcode_id"`
	Channel          string          `json:"channel"`
	AddressKey       string          `json:"address_key"`
}

type DepositPagesQuery struct {
	PerPage int             `json:"per_page"`
	Page    int             `json:"page"`
	Query   *DepositQuery   `json:"query"`
	OrderBy *DepositOrderBy `json:"order_by"`
}

type DepositQuery struct {
	Source         string     `json:"source"`
	To             *time.Time `json:"to"`
	From           *time.Time `json:"from"`
	CompanyID      int32      `json:"-"`
	Status         string     `json:"status"`
	Type           string     `json:"type"`
	Protocol       string     `json:"protocol"`
	ProcessingTo   *time.Time `json:"processing_to"`
	ProcessingFrom *time.Time `json:"processing_from"`
	Ids            []string   `json:"ids"`
}
type DepositOrderBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type DepositPagesResponse struct {
	Data  []*Deposit `json:"data"`
	Page  int        `json:"page"`
	Pages int        `json:"pages"`
	Total int        `json:"total"`
}

type CreateDeposit struct {
	Amount float32 `json:"amount"`
	Type   string  `json:"type"`
}

// Deposit is a structure manager all about Deposit
type DepositClient struct {
	client *Genesis
}

//Deposit - Instance de Deposit
func (g *Genesis) Deposit() *DepositClient {
	return &DepositClient{client: g}
}

//Create - create a new wallet
func (d *DepositClient) Create(req CreateDeposit) (*Deposit, *Error, error) {
	data, _ := json.Marshal(req)
	var response *Deposit
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/deposit", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//Create - create a new wallet
func (d *DepositClient) Get(token string) (*Deposit, *Error, error) {
	var response *Deposit
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "GET", fmt.Sprintf("payment/v1/api/deposit/%s", token), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (d *DepositClient) List(req DepositPagesQuery) (*DepositPagesResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *DepositPagesResponse
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/deposits", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
