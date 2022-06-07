package genesis

import (
	"encoding/json"
	"fmt"

	"github.com/medeirosfalante/bankly-go-sdk"
)

type PayBill struct {
	ID                int32   `json:"id"`
	Protocol          string  `json:"protocol"`
	BarCode           string  `json:"bar_code"`
	CreatedIn         int64   `json:"created_in"`
	UpdatedIn         int64   `json:"updated_in"`
	CompanyID         int32   `json:"company_id"`
	IntegrationID     string  `json:"integration_id"`
	Amount            float32 `json:"amount"`
	Status            string  `json:"status"`
	Object            string  `json:"object"`
	Digitableline     string  `json:"digitable_line"`
	AuthorizeID       string  `json:"authorize_id"`
	Type              string  `json:"type"`
	DueDate           string  `json:"dueDate"`
	TypeAuthorize     int32   `json:"type_authorize"`
	ErrorMessage      string  `json:"error_message"`
	DocumentPayer     string  `json:"document_payer"`
	DocumentRecipient string  `json:"document_recipient"`
	LiquidIn          int64   `json:"liquid_in"`
	SendIn            int64   `json:"send_in"`
	Description       string  `json:"description"`
	CorrelationID     string  `json:"correlation_id"`
}

type PayBillPagesQuery struct {
	PerPage int             `json:"per_page"`
	Page    int             `json:"page"`
	Query   *PayBillQuery   `json:"query"`
	OrderBy *PayBillOrderBy `json:"order_by"`
}

type PayBillQuery struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	To        int64  `json:"to"`
	From      int64  `json:"from"`
	CompanyID int32  `json:"-"`
}
type PayBillOrderBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type PayBillPagesResponse struct {
	Data  []*PayBill `json:"data"`
	Page  int        `json:"page"`
	Pages int        `json:"pages"`
	Total int        `json:"total"`
}

type PayBillBoleto struct {
	Digitableline string `json:"digitable_line"`
	BarCode       string `json:"bar_code"`
}

type CheckPayBillBoletoResponse struct {
	Amount        float32                      `json:"amount"`
	Balance       float32                      `json:"balance"`
	HaveBalance   bool                         `json:"have_balance"`
	BoletoObject  *bankly.ValidateBillResponse `json:"boleto_object"`
	CorrelationID string                       `json:"correlation_id"`
}

type ConfirmPayBill struct {
	Protocol    string          `json:"protocol"`
	Type        string          `json:"type"`
	PaymentCard *ConfirmPayBill `json:"paymentcard"`
}

type ConfirmPayBillCard struct {
}

type CreatePayBillResponse struct {
	PayBills *PayBill `json:"billing"`
}

type CreatePayBillCard struct {
	Protocol string `json:"protocol"`
	Link     string `json:"link"`
	DueDate  string `json:"due_date"`
}

// paybill is a structure manager all about paybill
type PayBillClient struct {
	client *Genesis
}

//paybill - Instance de paybill
func (g *Genesis) PayBill() *PayBillClient {
	return &PayBillClient{client: g}
}

//Create - create a
func (d *PayBillClient) Create(req PayBillBoleto) (*PayBill, *Error, error) {
	data, _ := json.Marshal(req)
	var response *PayBill
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/paybill/boleto", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//Create -  paybill
func (d *PayBillClient) Check(req PayBillBoleto) (*CheckPayBillBoletoResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *CheckPayBillBoletoResponse
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/paybill/boleto/check", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//Get - paybill details
func (d *PayBillClient) Get(token string) (*PayBill, *Error, error) {
	var response *PayBill
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "GET", fmt.Sprintf("payment/v1/api/paybill/%s", token), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//List - list alls paybills
func (d *PayBillClient) List(req PayBillPagesQuery) (*PayBillPagesResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *PayBillPagesResponse
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/paybills", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
