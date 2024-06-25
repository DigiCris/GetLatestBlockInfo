package blockchain

import (
    "fmt"
    "log"

    //"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockInfo() {/*
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
    if err != nil {
        log.Fatal(err)
    }

    block, err := client.BlockByNumber(nil, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Last block number: %d\n", block.Number())
    fmt.Printf("Last block hash: %x\n", block.Hash())
    fmt.Printf("Last block timestamp: %d\n", block.Time())
	*/
	fmt.Println("Hello World")
}