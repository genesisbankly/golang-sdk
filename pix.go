package genesis

import (
	"encoding/json"
	"fmt"
	"time"
)

type PixKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PixKeyHolder struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type PixKeyResponse struct {
	EndToEndId    string        `json:"endToEndId"`
	AddressingKey *PixKey       `json:"addressingKey"`
	Status        string        `json:"status"`
	CreatedAt     *time.Time    `json:"createdAt"`
	OwnedAt       *time.Time    `json:"ownedAt"`
	Holder        *PixKeyHolder `json:"holder,omitempty"`
}

// Deposit is a structure manager all about Deposit
type PixClient struct {
	client *Genesis
}

type CreateBrcodeRequest struct {
	PixId       int32   `json:"pix_id"`
	Amount      float32 `json:"amount,omitempty"`
	Description string  `json:"description"`
}

type Brcode struct {
	ID          int32     `json:"id"`
	Active      bool      `json:"active"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Deleted     bool      `json:"-"`
	CompanyID   int32     `json:"company_id"`
	EndToEndId  string    `json:"end_to_end_id"`
	AccountID   int32     `json:"account_id"`
	Protocol    string    `json:"protocol"`
	Description string    `json:"description"`
}

//Deposit - Instance de Deposit
func (g *Genesis) Pix() *PixClient {
	return &PixClient{client: g}
}

func (d *PixClient) GetKey(addressingKeyValue string) (*PixKeyResponse, *Error, error) {
	var response *PixKeyResponse
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "GET", fmt.Sprintf("payment/v1/api/pix/key/%s", addressingKeyValue), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (d *PixClient) CreateBrcode(req CreateBrcodeRequest) (*Brcode, *Error, error) {
	data, _ := json.Marshal(req)
	var response *Brcode
	responseToken, err := d.client.RequestToken()
	if err != nil {
		return nil, nil, err
	}
	err, errAPI := d.client.Request(responseToken.AccessToken, "POST", "payment/v1/api/brcode", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
