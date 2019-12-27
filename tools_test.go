package CocosGoSDKTool

import (
	sdk "CocosSDK"
	"encoding/json"
	"testing"
)

const MAIN_NET_FAUCET = "https://faucet.cocosbcx.net/api/v1/accounts"
const TEST_FAUCET = "https://test-faucet.cocosbcx.net/api/v1/accounts"

func TestInitSdk(t *testing.T) {
	sdk.InitSDK("test.cocosbcx.net", true)
	//sdk.Wallet.AddAccountByPrivateKey("5KBWr3cR5rdekfLaHAnaEFsoc8uGqHbG5CDGSdZqkCuy1ajSXzJ","1234")
	//sdk.Wallet.ImportAccount("ggggxxx", "12345678")
	//sdk.Wallet.SetDefaultAccount("oooo00", "1234")
	//t.Log(sdk.Wallet.CreateAccount("o-.-o","ximenmaohui"))
	//t.Log(rpc.GetDynamicGlobalProperties())
}

func TestTxsForAddress(t *testing.T) {
	txs, err := TxsForAddress("bitpie.com44",1, "4b1c4e80232a7f3d904bfacb6bcb045620e3eeae54df0f973e6e09f339249ab0")
	t.Log(err)//, 10, "263c47271171a5f99c839475f232a742074f848ddaba558e1b151106bf8dfbd1")
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}
