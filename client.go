package genesis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type Genesis struct {
	client       *http.Client
	ClientID     string
	ClientSecret string
	Env          string
	Token        string
}

type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewClient(ClientID, ClientSecret, env string) *Genesis {
	genesis := &Genesis{
		client:       &http.Client{Timeout: 60 * time.Second},
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Env:          env,
	}
	return genesis

}

func (genesis *Genesis) Request(token, method, action string, body []byte, out interface{}) (error, *Error) {
	url := genesis.devProd()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	_, err, errBody := genesis.RequestMaster(req, &out)
	return err, errBody
}

func (genesis *Genesis) RequestMaster(req *http.Request, out interface{}) ([]byte, error, *Error) {
	res, err := genesis.client.Do(req)
	if err != nil {
		return nil, err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 202 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return bodyResponse, err, nil
		}
		errAPI.Body = string(bodyResponse)
		return bodyResponse, nil, &errAPI
	}
	if out != nil {
		err = json.Unmarshal(bodyResponse, out)
		if err != nil {
			return bodyResponse, err, nil
		}
	}

	return bodyResponse, nil, nil
}

func (Genesis *Genesis) devProd() string {
	if Genesis.Env == "production" {
		return "https://api.genesisapp.cloud"
	}
	return "https://api.sandbox.genesisapp.cloud"
}

func (genesis *Genesis) RequestToken() (*TokenResponse, error) {
	var tokenResponse TokenResponse
	url := genesis.devProd()

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("client_secret", genesis.ClientSecret)
	writer.WriteField("grant_type", "all")
	writer.WriteField("client_id", genesis.ClientID)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", url, "customer/v1/connect/token")
	fmt.Printf("payload %s\n", payload.String())
	req, err := http.NewRequest("POST", endpoint, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	fmt.Printf("response %s\n", string(bodyResponse))
	if res.StatusCode > 202 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyResponse))
	}
	err = json.Unmarshal(bodyResponse, &tokenResponse)
	if err != nil {
		return nil, err
	}
	genesis.Token = tokenResponse.AccessToken
	return &tokenResponse, nil
}
