package blockchain_test

import (
	"github.com/qedus/blockchain"
	"net/http"
	"testing"
)

var (
	bc = blockchain.New(http.DefaultClient)
)

func init() {
	bc.GUID = "[required wallet GUID]"
	bc.Password = "[required wallet password]"
	bc.SecondPassword = "[optional second wallet password]"
	bc.APICode = "[optional API code]"
}

func TestNewAddress(t *testing.T) {
	na := &blockchain.NewAddress{Label: "testlabel"}
	if err := bc.Request(na); err != nil {
		t.Fatal(err)
	}

	if na.Label != "testlabel" {
		t.Fatalf("not correct label in new address [%s]", na.Label)
	}

	if na.Address == "" {
		t.Fatal("no new address")
	}
}

func TestSendPayment(t *testing.T) {
	sp := &blockchain.SendPayment{
		ToAddress:"1BTCorgHwCg6u2YSAWKgS17qUad6kHmtQW",
		Amount:1}
	if err := bc.Request(sp); err != nil {
		t.Fatal(err)
	}
}
