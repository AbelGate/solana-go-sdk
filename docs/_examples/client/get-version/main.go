package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AbelGate/solana-go-sdk/client"
	"github.com/AbelGate/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	resp, err := c.GetVersion(context.TODO())
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
	}

	fmt.Println("version", resp.SolanaCore)
}
