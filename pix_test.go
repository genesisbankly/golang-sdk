package genesis_test

import (
	"os"
	"testing"

	genesis "github.com/genesisbankly/golang-sdk"
	"github.com/joho/godotenv"
)

func TestGetKey(t *testing.T) {
	godotenv.Load(".env.test")
	client := genesis.NewClient(os.Getenv("GENESIS_CLIENT_ID"), os.Getenv("GENESIS_CLIENT_SECRET"), os.Getenv("ENV"))
	pixResponse, errAPI, err := client.Pix().GetKey("03602763501")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if pixResponse == nil {
		t.Error("payResponse is null")
		return
	}

}

func TestCreateBrcode(t *testing.T) {
	godotenv.Load(".env.test")
	client := genesis.NewClient(os.Getenv("GENESIS_CLIENT_ID"), os.Getenv("GENESIS_CLIENT_SECRET"), os.Getenv("ENV"))
	pixResponse, errAPI, err := client.Pix().CreateBrcode(genesis.CreateBrcodeRequest{
		PixId:       1,
		Description: "10",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if pixResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
