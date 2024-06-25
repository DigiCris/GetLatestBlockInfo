package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockInfo(rpc string) {
    ctx := context.Background() // Crear un contexto de fondo
    client, err := ethclient.Dial(rpc)
    if err != nil {
        log.Fatal(err)
    }

    block, err := client.BlockByNumber(ctx, nil) // Pasar el contexto aquí
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Last block number: %d\n", block.Number())
    fmt.Printf("Last block hash: %x\n", block.Hash())
    fmt.Printf("Last block timestamp: %d\n", block.Time())
}
