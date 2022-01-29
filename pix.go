package genesis

import (
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
