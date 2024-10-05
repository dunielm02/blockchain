package main

import (
	"github.com/dunielm02/blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "Basic blockchain in golang.",
	Long:  `Basic blockchain in golang.`,
}

var bc *blockchain.Blockchain

func main() {
	bc = blockchain.NewBlockchain()

	rootCmd.AddCommand(addBlock)
	rootCmd.AddCommand(printChain)

	rootCmd.Execute()
}
