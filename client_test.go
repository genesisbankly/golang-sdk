package genesis_test

import (
	"os"
	"testing"

	genesis "github.com/genesisbankly/golang-sdk"
	"github.com/joho/godotenv"
)

func TestRequesttoken(t *testing.T) {
	godotenv.Load(".env.test")
	client := genesis.NewClient(os.Getenv("GENESIS_CLIENT_ID"), os.Getenv("GENESIS_CLIENT_SECRET"), os.Getenv("ENV"))

	response, err := client.RequestToken()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(response.AccessToken) <= 0 {
		t.Errorf("AccessToken is invalid")
	}

}
