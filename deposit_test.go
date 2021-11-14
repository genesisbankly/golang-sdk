package genesis_test

import (
	"os"
	"testing"

	genesis "github.com/genesisbankly/golang-sdk"
	"github.com/joho/godotenv"
)

func TestListDeposits(t *testing.T) {
	godotenv.Load(".env.test")
	client := genesis.NewClient(os.Getenv("GENESIS_CLIENT_ID"), os.Getenv("GENESIS_CLIENT_SECRET"), os.Getenv("ENV"))
	depositPagesQuery := genesis.DepositPagesQuery{
		PerPage: 10,
		Page:    1,
	}
	depositPagesResponse, errAPI, err := client.Deposit().List(depositPagesQuery)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	t.Error("payResponse is null")
	if depositPagesResponse == nil {
		t.Error("payResponse is null")
		return
	}

}

func TestCreate(t *testing.T) {
	godotenv.Load(".env.test")
	client := genesis.NewClient(os.Getenv("GENESIS_CLIENT_ID"), os.Getenv("GENESIS_CLIENT_SECRET"), os.Getenv("ENV"))
	depositPagesResponse, errAPI, err := client.Deposit().Create(genesis.CreateDeposit{
		Amount: 10,
		Type:   "brcode",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	t.Error("payResponse is null")
	if depositPagesResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
