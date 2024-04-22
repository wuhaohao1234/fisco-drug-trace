package sdk

import (
	"context"
	"encoding/hex"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

var Contract *DrugTraceSession

func InitContract() {
	// 连接fisco 并初始化合约
	privateKey, _ := hex.DecodeString("d4ed1c42b66e4e7040bed37bd041f89a4ce4b9e899a3b8631949ea626c135161")
	// disable ssl of node rpc
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", DisableSsl: false,
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./sdk/ca.crt", TLSKeyFile: "./sdk/sdk.key", TLSCertFile: "./sdk/sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x1d38f5d0c8c1ae7ed63a2d0ec905b9e9a17e70cf")
	instance, err := NewDrugTrace(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	Contract = &DrugTraceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
}
