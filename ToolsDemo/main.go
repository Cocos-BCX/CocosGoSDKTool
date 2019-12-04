// ToolsDemo project main.go
package main

import (
	"fmt"
	"log"

	tools "CocosGoSDKTool"
	sdk "CocosSDK"
)

func main() {
	sdk.InitSDK("test.cocosbcx.net", 80, false)
	sdk.Wallet.ImportAccount("gggg1", "12345678")
	sdk.Wallet.SetDefaultAccount("gggg1", "12345678")
	log.Println(tools.CreateAccount("sqctwe123", "0x02703d7df82c35218fbc459f49f3ae918c29fc68665f4689b8248808bbf79bddc2"))
	txs, _ := tools.TxsForAddress("gggg1")
	fmt.Println(txs)
	hash := tools.UnsignedTxHash("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad008bcc75d6e5f01f4cf5d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb0326e614f3b308beaa01081355d8f9325a76997c83b2dfb17652e0000")
	fmt.Println(hash)
}
