package CocosGoSDKTool

import (
	sdk "CocosSDK"
	"CocosSDK/rpc"
	"encoding/json"
	"testing"
)

func TestInitSdk(t *testing.T) {
	sdk.InitSDK("test.cocosbcx.net", true)
	sdk.Wallet.ImportAccount("ggggxxx", "12345678")
	sdk.Wallet.SetDefaultAccount("ggggxxx", "12345678")
	t.Log(rpc.GetDynamicGlobalProperties())
}
func TestGetTransaction(t *testing.T) {
	tx, err := GetTransaction("4432a6f92b95ade128f52e378f376eec87bfa50230c26584974554d1ab730c66")
	t.Log(err)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}
func TestDeserialize(t *testing.T) {
	var hex_str string = "1ae3653a3105800f5722c5bda2b55530d0e9e8654314e2f3dc6d2b010da641c5cd27a29b23cab399e75d010063490000000000003000000000000000a0860100000000000000000000000000010103679b27673fea2846434ca659ffe65ea8c0ec6e751aa9c326bbfbfcf8affa673b0354702c8f8a62a9b0ec52ff6ef3d439cad11ddec8a105af00deab548788ab56fd99f3124df33a81401047e1947d6e104ec19ac290bb159e71960000"
	tx, err := Deserialize(hex_str)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestDeserializeTransactions(t *testing.T) {
	sign_tx, _ := DeserializeTransactions("1ae3653a3105800f5722c5bda2b55530d0e9e8654314e2f3dc6d2b010da641c5cd27a29b23cab299e75d010063490000000000003000000000000000a0860100000000000000000000000000010004787878780000")
	byte_s, err := json.Marshal(sign_tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestTxsForAddress(t *testing.T) {
	txs, err := TxsForAddress("ggggxxx") //, 10, "263c47271171a5f99c839475f232a742074f848ddaba558e1b151106bf8dfbd1")
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestPuk2Addr(t *testing.T) {
	t.Log(PublicToAddress("0x02703d7df82c35218fbc459f49f3ae918c29fc68665f4689b8248808bbf79bddc2"))
}

func TestAddr2Puk(t *testing.T) {
	t.Log(AddressToPublic("ggggxxx"))
}

func TestGetBlockCount(t *testing.T) {
	t.Log(Getblockcount())
}

/*
func TestGetrawmempool(t *testing.T) {
	t.Log(Getrawmempool())
}*/

func TestGetblocktxs(t *testing.T) {
	txs, err := Getblocktxs(77559)
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestBalanceForAddress(t *testing.T) {
	balances := BalanceForAddress("ggggxxx")
	byte_s, err := json.Marshal(balances)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestBalanceForAddressForCoinCode(t *testing.T) {
	balances := BalanceForAddressForCoinCode("test1", "COCOS")
	byte_s, err := json.Marshal(balances)
	if err == nil {
		t.Log(string(byte_s))
	}
	balances = BalanceForAddressForCoinCode("test1", "1.3.1")
	byte_s, err = json.Marshal(balances)
	if err == nil {
		t.Log(string(byte_s))
	}
}
func TestSignTransaction(t *testing.T) {
	tx, err := SignTransaction("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0486711d2c551899dc85d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb032b4d412ed0c8e38561077883f0dfb4c3f8e1068c92ef3e9653f0000",
		[]string{"202c76ab413de66315922a95c65b0dc77073bf1f9a7e809b0aa51db9f1592e359c2de34ed115c039d356ca573e0d4dc818a258acfc0af48c44c6e4c8d2c9d57508"})

	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestBuildTransaction(t *testing.T) {
	hex_str, err := BuildTransaction("ggggxxx", "test1", 1, "COCOS")
	t.Log(err)
	t.Log(hex_str)
}

func TestUnsignedTxHash(t *testing.T) {
	hash := UnsignedTxHash("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad008bcc75d6e5f01f4cf5d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb0326e614f3b308beaa01081355d8f9325a76997c83b2dfb17652e0000")
	if hash == "263c47271171a5f99c839475f232a742074f848ddaba558e1b151106bf8dfbd1" {
		t.Log("Test Unsigned Tx Hash success!!")
	} else {
		t.Error("Test Unsigned Tx Hash Error!")
	}
}

func TestCreateAccount(t *testing.T) {
	t.Log(CreateAccount("sqctccc123", "0x02703d34f82c35218fbc459f49f3ae918c29fc68665f4689b8248808bbf79bddc2"))

}
